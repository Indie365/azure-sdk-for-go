package storage

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// OverrideHeaders defines overridable response heaedrs in
// a request using a SAS URI.
// See https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas
type OverrideHeaders struct {
	CacheControl       string
	ContentDisposition string
	ContentEncoding    string
	ContentLanguage    string
	ContentType        string
}

// BlobSASOptions are options to construct a blob SAS
// URI.
// See https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas
type BlobSASOptions struct {
	BlobSASPermissions
	OverrideHeaders
	SASOptions
}

// SASOptions includes options used by SAS URIs for diffrenet
// services and resources.
type SASOptions struct {
	APIVersion string
	Start      time.Time
	Expiry     time.Time
	IP         string
	UseHTTPS   bool
	Identifier string
}

// BlobSASPermissions includes the available permissions for
// a blob SAS URI.
type BlobSASPermissions struct {
	Read   bool
	Add    bool
	Create bool
	Write  bool
	Delete bool
}

// GetSASURI creates an URL to the blob which contains the Shared
// Access Signature with the specified options.
//
// See https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas
func (b *Blob) GetSASURI(options BlobSASOptions) (string, error) {
	uri := b.GetURL()
	signedResource := "b"
	canonicalizedResource, err := b.Container.bsc.client.buildCanonicalizedResource(uri, b.Container.bsc.auth)
	if err != nil {
		return "", err
	}

	// build permissions string
	permissions := ""
	if options.Read {
		permissions += "r"
	}
	if options.Add {
		permissions += "a"
	}
	if options.Create {
		permissions += "c"
	}
	if options.Write {
		permissions += "w"
	}
	if options.Delete {
		permissions += "d"
	}

	return b.Container.bsc.client.commonSASURI(options.SASOptions, uri, permissions, canonicalizedResource, signedResource, options.OverrideHeaders)
}

func (c *Client) commonSASURI(options SASOptions, uri, permissions, canonicalizedResource, signedResource string, headers OverrideHeaders) (string, error) {
	start := ""
	if options.Start != (time.Time{}) {
		start = options.Start.UTC().Format(time.RFC3339)
	}

	expiry := options.Expiry.UTC().Format(time.RFC3339)

	// We need to replace + with %2b first to avoid being treated as a space (which is correct for query strings, but not the path component).
	canonicalizedResource = strings.Replace(canonicalizedResource, "+", "%2b", -1)
	canonicalizedResource, err := url.QueryUnescape(canonicalizedResource)
	if err != nil {
		return "", err
	}

	protocols := "https,http"
	if options.UseHTTPS {
		protocols = "https"
	}
	stringToSign, err := blobSASStringToSign(permissions, start, expiry, canonicalizedResource, options.Identifier, options.IP, protocols, c.apiVersion, headers)
	if err != nil {
		return "", err
	}

	sig := c.computeHmac256(stringToSign)
	sasParams := url.Values{
		"sv":  {c.apiVersion},
		"se":  {expiry},
		"sr":  {signedResource},
		"sp":  {permissions},
		"sig": {sig},
	}

	addQueryParameter(sasParams, "st", start)
	addQueryParameter(sasParams, "si", options.Identifier)

	if c.apiVersion >= "2015-04-05" {
		sasParams.Add("spr", protocols)
		addQueryParameter(sasParams, "sip", options.IP)
	}

	// Add override response hedaers
	addQueryParameter(sasParams, "rscc", headers.CacheControl)
	addQueryParameter(sasParams, "rscd", headers.ContentDisposition)
	addQueryParameter(sasParams, "rsce", headers.ContentEncoding)
	addQueryParameter(sasParams, "rscl", headers.ContentLanguage)
	addQueryParameter(sasParams, "rsct", headers.ContentType)

	sasURL, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	sasURL.RawQuery = sasParams.Encode()
	return sasURL.String(), nil
}

func addQueryParameter(query url.Values, key, value string) url.Values {
	if value != "" {
		query.Add(key, value)
	}
	return query
}

func blobSASStringToSign(signedPermissions, signedStart, signedExpiry, canonicalizedResource, signedIdentifier, signedIP, protocols, signedVersion string, headers OverrideHeaders) (string, error) {
	rscc := headers.CacheControl
	rscd := headers.ContentDisposition
	rsce := headers.ContentEncoding
	rscl := headers.ContentLanguage
	rsct := headers.ContentType

	if signedVersion >= "2015-02-21" {
		canonicalizedResource = "/blob" + canonicalizedResource
	}

	// https://msdn.microsoft.com/en-us/library/azure/dn140255.aspx#Anchor_12
	if signedVersion >= "2015-04-05" {
		return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s", signedPermissions, signedStart, signedExpiry, canonicalizedResource, signedIdentifier, signedIP, protocols, signedVersion, rscc, rscd, rsce, rscl, rsct), nil
	}

	// reference: http://msdn.microsoft.com/en-us/library/azure/dn140255.aspx
	if signedVersion >= "2013-08-15" {
		return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s", signedPermissions, signedStart, signedExpiry, canonicalizedResource, signedIdentifier, signedVersion, rscc, rscd, rsce, rscl, rsct), nil
	}

	return "", errors.New("storage: not implemented SAS for versions earlier than 2013-08-15")
}
