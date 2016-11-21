package storage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// TableServiceClient contains operations for Microsoft Azure Table Storage
// Service.
type TableServiceClient struct {
	client Client
}

// AzureTable is the typedef of the Azure Table name
type AzureTable string

const (
	tablesURIPath = "/Tables"
)

type createTableRequest struct {
	TableName string `json:"TableName"`
}

// TableAccessPolicyDetails are used for SETTING table policies
type TableAccessPolicyDetails struct {
	ID         string
	StartTime  time.Time
	ExpiryTime time.Time
	CanRead    bool
	CanAppend  bool
	CanUpdate  bool
	CanDelete  bool
}

func pathForTable(table AzureTable) string { return fmt.Sprintf("%s", table) }

func (c *TableServiceClient) getStandardHeaders() map[string]string {
	return map[string]string{
		"x-ms-version":   "2015-02-21",
		"x-ms-date":      currentTimeRfc1123Formatted(),
		"Accept":         "application/json;odata=nometadata",
		"Accept-Charset": "UTF-8",
		"Content-Type":   "application/json",
	}
}

// QueryTables returns the tables created in the
// *TableServiceClient storage account.
func (c *TableServiceClient) QueryTables() ([]AzureTable, error) {
	uri := c.client.getEndpoint(tableServiceName, tablesURIPath, url.Values{})

	headers := c.getStandardHeaders()
	headers["Content-Length"] = "0"

	resp, err := c.client.execTable("GET", uri, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.body.Close()

	if err := checkRespCode(resp.statusCode, []int{http.StatusOK}); err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.body)

	var respArray queryTablesResponse
	if err := json.Unmarshal(buf.Bytes(), &respArray); err != nil {
		return nil, err
	}

	s := make([]AzureTable, len(respArray.TableName))
	for i, elem := range respArray.TableName {
		s[i] = AzureTable(elem.TableName)
	}

	return s, nil
}

// CreateTable creates the table given the specific
// name. This function fails if the name is not compliant
// with the specification or the tables already exists.
func (c *TableServiceClient) CreateTable(table AzureTable) error {
	uri := c.client.getEndpoint(tableServiceName, tablesURIPath, url.Values{})

	headers := c.getStandardHeaders()

	req := createTableRequest{TableName: string(table)}
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return err
	}

	headers["Content-Length"] = fmt.Sprintf("%d", buf.Len())

	resp, err := c.client.execTable("POST", uri, headers, buf)

	if err != nil {
		return err
	}
	defer resp.body.Close()

	if err := checkRespCode(resp.statusCode, []int{http.StatusCreated}); err != nil {
		return err
	}

	return nil
}

// DeleteTable deletes the table given the specific
// name. This function fails if the table is not present.
// Be advised: DeleteTable deletes all the entries
// that may be present.
func (c *TableServiceClient) DeleteTable(table AzureTable) error {
	uri := c.client.getEndpoint(tableServiceName, tablesURIPath, url.Values{})
	uri += fmt.Sprintf("('%s')", string(table))

	headers := c.getStandardHeaders()

	headers["Content-Length"] = "0"

	resp, err := c.client.execTable("DELETE", uri, headers, nil)

	if err != nil {
		return err
	}
	defer resp.body.Close()

	if err := checkRespCode(resp.statusCode, []int{http.StatusNoContent}); err != nil {
		return err

	}
	return nil
}

// SetTablePermissions sets up table ACL permissinos.
func (c *TableServiceClient) SetTablePermissions(table AzureTable, accessPolicy TableAccessPolicyDetails, timeout int) (err error) {
	params := url.Values{
		"comp": {"acl"},
	}

	if timeout > 0 {
		params.Add("timeout", strconv.Itoa(timeout))
	}

	uri := c.client.getEndpoint(tableServiceName, string(table), params)
	headers := c.client.getStandardHeaders()

	var permissions = generateTablePermissions(accessPolicy)

	// generate the XML for the SharedAccessSignature if required.
	accessPolicyXML, err := generateAccessPolicy(accessPolicy.ID, accessPolicy.StartTime,
		accessPolicy.ExpiryTime,
		permissions)
	if err != nil {
		return err
	}

	var resp *storageResponse
	if accessPolicyXML != "" {
		headers["Content-Length"] = strconv.Itoa(len(accessPolicyXML))
		resp, err = c.client.exec("PUT", uri, headers, strings.NewReader(accessPolicyXML))
	} else {
		resp, err = c.client.exec("PUT", uri, headers, nil)
	}

	if err != nil {
		return err
	}

	if resp != nil {
		defer func() {
			err = resp.body.Close()
		}()

		if resp.statusCode != http.StatusOK {
			return errors.New("Unable to set permissions")
		}
	}
	return nil
}

// GetTablePermissions gets the table ACL permissions
func (c *TableServiceClient) GetTablePermissions(table AzureTable, timeout int) (permissionResponse *AccessPolicy, err error) {
	params := url.Values{"comp": {"acl"}}

	if timeout > 0 {
		params.Add("timeout", strconv.Itoa(timeout))
	}

	uri := c.client.getEndpoint(tableServiceName, string(table), params)
	log.Printf("get url %s", uri)

	headers := c.client.getStandardHeaders()

	resp, err := c.client.execTable("GET", uri, headers, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.body.Close()
	}()

	var out AccessPolicy
	err = xmlUnmarshal(resp.body, &out.SignedIdentifiersList)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func generateTablePermissions(accessPolicy TableAccessPolicyDetails) (permissions string) {
	// generate the permissions string (raud).
	// still want the end user API to have bool flags.
	permissions = ""

	if accessPolicy.CanRead {
		permissions += "r"
	}

	if accessPolicy.CanAppend {
		permissions += "a"
	}

	if accessPolicy.CanUpdate {
		permissions += "u"
	}

	if accessPolicy.CanDelete {
		permissions += "d"
	}
	return permissions
}
