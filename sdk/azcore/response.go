//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcore

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Payload reads and returns the response body or an error.
// On a successful read, the response body is cached.
func Payload(resp *http.Response) ([]byte, error) {
	// r.Body won't be a nopClosingBytesReader if downloading was skipped
	if buf, ok := resp.Body.(*nopClosingBytesReader); ok {
		return buf.Bytes(), nil
	}
	bytesBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	resp.Body = &nopClosingBytesReader{s: bytesBody, i: 0}
	return bytesBody, nil
}

// HasStatusCode returns true if the Response's status code is one of the specified values.
func HasStatusCode(resp *http.Response, statusCodes ...int) bool {
	if resp == nil {
		return false
	}
	for _, sc := range statusCodes {
		if resp.StatusCode == sc {
			return true
		}
	}
	return false
}

// UnmarshalAsByteArray will base-64 decode the received payload and place the result into the value pointed to by v.
func UnmarshalAsByteArray(resp *http.Response, v *[]byte, format Base64Encoding) error {
	p, err := Payload(resp)
	if err != nil {
		return err
	}
	return DecodeByteArray(string(p), v, format)
}

// UnmarshalAsJSON calls json.Unmarshal() to unmarshal the received payload into the value pointed to by v.
func UnmarshalAsJSON(resp *http.Response, v interface{}) error {
	payload, err := Payload(resp)
	if err != nil {
		return err
	}
	// TODO: verify early exit is correct
	if len(payload) == 0 {
		return nil
	}
	err = removeBOM(resp)
	if err != nil {
		return err
	}
	err = json.Unmarshal(payload, v)
	if err != nil {
		err = fmt.Errorf("unmarshalling type %T: %s", v, err)
	}
	return err
}

// UnmarshalAsXML calls xml.Unmarshal() to unmarshal the received payload into the value pointed to by v.
func UnmarshalAsXML(resp *http.Response, v interface{}) error {
	payload, err := Payload(resp)
	if err != nil {
		return err
	}
	// TODO: verify early exit is correct
	if len(payload) == 0 {
		return nil
	}
	err = removeBOM(resp)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(payload, v)
	if err != nil {
		err = fmt.Errorf("unmarshalling type %T: %s", v, err)
	}
	return err
}

// Drain reads the response body to completion then closes it.  The bytes read are discarded.
func Drain(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}

// removeBOM removes any byte-order mark prefix from the payload if present.
func removeBOM(resp *http.Response) error {
	payload, err := Payload(resp)
	if err != nil {
		return err
	}
	// UTF8
	trimmed := bytes.TrimPrefix(payload, []byte("\xef\xbb\xbf"))
	if len(trimmed) < len(payload) {
		resp.Body.(*nopClosingBytesReader).Set(trimmed)
	}
	return nil
}

// writes to a buffer, used for logging purposes
func writeBody(resp *http.Response, b *bytes.Buffer) error {
	ct := resp.Header.Get(headerContentType)
	if ct == "" {
		fmt.Fprint(b, "   Response contained no body\n")
		return nil
	} else if !shouldLogBody(b, ct) {
		return nil
	}
	body, err := Payload(resp)
	if err != nil {
		fmt.Fprintf(b, "   Failed to read response body: %s\n", err.Error())
		return err
	}
	if len(body) > 0 {
		logBody(b, body)
	} else {
		fmt.Fprint(b, "   Response contained no body\n")
	}
	return nil
}

// RetryAfter returns non-zero if the response contains a Retry-After header value.
func RetryAfter(resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}
	ra := resp.Header.Get(headerRetryAfter)
	if ra == "" {
		return 0
	}
	// retry-after values are expressed in either number of
	// seconds or an HTTP-date indicating when to try again
	if retryAfter, _ := strconv.Atoi(ra); retryAfter > 0 {
		return time.Duration(retryAfter) * time.Second
	} else if t, err := time.Parse(time.RFC1123, ra); err == nil {
		return time.Until(t)
	}
	return 0
}

// DecodeByteArray will base-64 decode the provided string into v.
func DecodeByteArray(s string, v *[]byte, format Base64Encoding) error {
	if len(s) == 0 {
		return nil
	}
	payload := string(s)
	if payload[0] == '"' {
		// remove surrounding quotes
		payload = payload[1 : len(payload)-1]
	}
	switch format {
	case Base64StdFormat:
		decoded, err := base64.StdEncoding.DecodeString(payload)
		if err == nil {
			*v = decoded
			return nil
		}
		return err
	case Base64URLFormat:
		// use raw encoding as URL format should not contain any '=' characters
		decoded, err := base64.RawURLEncoding.DecodeString(payload)
		if err == nil {
			*v = decoded
			return nil
		}
		return err
	default:
		return fmt.Errorf("unrecognized byte array format: %d", format)
	}
}

// writeRequestWithResponse appends a formatted HTTP request into a Buffer. If request and/or err are
// not nil, then these are also written into the Buffer.
func writeRequestWithResponse(b *bytes.Buffer, request *Request, resp *http.Response, err error) {
	// Write the request into the buffer.
	fmt.Fprint(b, "   "+request.Method+" "+request.URL.String()+"\n")
	writeHeader(b, request.Header)
	if resp != nil {
		fmt.Fprintln(b, "   --------------------------------------------------------------------------------")
		fmt.Fprint(b, "   RESPONSE Status: "+resp.Status+"\n")
		writeHeader(b, resp.Header)
	}
	if err != nil {
		fmt.Fprintln(b, "   --------------------------------------------------------------------------------")
		fmt.Fprint(b, "   ERROR:\n"+err.Error()+"\n")
	}
}

// formatHeaders appends an HTTP request's or response's header into a Buffer.
func writeHeader(b *bytes.Buffer, header http.Header) {
	if len(header) == 0 {
		b.WriteString("   (no headers)\n")
		return
	}
	keys := make([]string, 0, len(header))
	// Alphabetize the headers
	for k := range header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		// Redact the value of any Authorization header to prevent security information from persisting in logs
		value := interface{}("REDACTED")
		if !strings.EqualFold(k, "Authorization") {
			value = header[k]
		}
		fmt.Fprintf(b, "   %s: %+v\n", k, value)
	}
}
