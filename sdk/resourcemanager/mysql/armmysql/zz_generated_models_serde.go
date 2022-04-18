//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AdvisorsResultList.
func (a AdvisorsResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ConfigurationListResult.
func (c ConfigurationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DatabaseListResult.
func (d DatabaseListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type FirewallRuleListResult.
func (f FirewallRuleListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LogFileListResult.
func (l LogFileListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LogFileProperties.
func (l LogFileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdTime", l.CreatedTime)
	populateTimeRFC3339(objectMap, "lastModifiedTime", l.LastModifiedTime)
	populate(objectMap, "sizeInKB", l.SizeInKB)
	populate(objectMap, "type", l.Type)
	populate(objectMap, "url", l.URL)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LogFileProperties.
func (l *LogFileProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdTime":
			err = unpopulateTimeRFC3339(val, &l.CreatedTime)
			delete(rawMsg, key)
		case "lastModifiedTime":
			err = unpopulateTimeRFC3339(val, &l.LastModifiedTime)
			delete(rawMsg, key)
		case "sizeInKB":
			err = unpopulate(val, &l.SizeInKB)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, &l.Type)
			delete(rawMsg, key)
		case "url":
			err = unpopulate(val, &l.URL)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	populate(objectMap, "properties", o.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PerformanceTierListResult.
func (p PerformanceTierListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PerformanceTierProperties.
func (p PerformanceTierProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "maxBackupRetentionDays", p.MaxBackupRetentionDays)
	populate(objectMap, "maxLargeStorageMB", p.MaxLargeStorageMB)
	populate(objectMap, "maxStorageMB", p.MaxStorageMB)
	populate(objectMap, "minBackupRetentionDays", p.MinBackupRetentionDays)
	populate(objectMap, "minLargeStorageMB", p.MinLargeStorageMB)
	populate(objectMap, "minStorageMB", p.MinStorageMB)
	populate(objectMap, "serviceLevelObjectives", p.ServiceLevelObjectives)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceListResult.
func (p PrivateLinkResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", p.NextLink)
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type QueryStatisticProperties.
func (q QueryStatisticProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregationFunction", q.AggregationFunction)
	populate(objectMap, "databaseNames", q.DatabaseNames)
	populateTimeRFC3339(objectMap, "endTime", q.EndTime)
	populate(objectMap, "metricDisplayName", q.MetricDisplayName)
	populate(objectMap, "metricName", q.MetricName)
	populate(objectMap, "metricValue", q.MetricValue)
	populate(objectMap, "metricValueUnit", q.MetricValueUnit)
	populate(objectMap, "queryExecutionCount", q.QueryExecutionCount)
	populate(objectMap, "queryId", q.QueryID)
	populateTimeRFC3339(objectMap, "startTime", q.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryStatisticProperties.
func (q *QueryStatisticProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aggregationFunction":
			err = unpopulate(val, &q.AggregationFunction)
			delete(rawMsg, key)
		case "databaseNames":
			err = unpopulate(val, &q.DatabaseNames)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &q.EndTime)
			delete(rawMsg, key)
		case "metricDisplayName":
			err = unpopulate(val, &q.MetricDisplayName)
			delete(rawMsg, key)
		case "metricName":
			err = unpopulate(val, &q.MetricName)
			delete(rawMsg, key)
		case "metricValue":
			err = unpopulate(val, &q.MetricValue)
			delete(rawMsg, key)
		case "metricValueUnit":
			err = unpopulate(val, &q.MetricValueUnit)
			delete(rawMsg, key)
		case "queryExecutionCount":
			err = unpopulate(val, &q.QueryExecutionCount)
			delete(rawMsg, key)
		case "queryId":
			err = unpopulate(val, &q.QueryID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &q.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type QueryTextsResultList.
func (q QueryTextsResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", q.NextLink)
	populate(objectMap, "value", q.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RecommendationActionProperties.
func (r RecommendationActionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "actionId", r.ActionID)
	populate(objectMap, "advisorName", r.AdvisorName)
	populateTimeRFC3339(objectMap, "createdTime", r.CreatedTime)
	populate(objectMap, "details", r.Details)
	populateTimeRFC3339(objectMap, "expirationTime", r.ExpirationTime)
	populate(objectMap, "reason", r.Reason)
	populate(objectMap, "recommendationType", r.RecommendationType)
	populate(objectMap, "sessionId", r.SessionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecommendationActionProperties.
func (r *RecommendationActionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "actionId":
			err = unpopulate(val, &r.ActionID)
			delete(rawMsg, key)
		case "advisorName":
			err = unpopulate(val, &r.AdvisorName)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateTimeRFC3339(val, &r.CreatedTime)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, &r.Details)
			delete(rawMsg, key)
		case "expirationTime":
			err = unpopulateTimeRFC3339(val, &r.ExpirationTime)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &r.Reason)
			delete(rawMsg, key)
		case "recommendationType":
			err = unpopulate(val, &r.RecommendationType)
			delete(rawMsg, key)
		case "sessionId":
			err = unpopulate(val, &r.SessionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RecommendationActionsResultList.
func (r RecommendationActionsResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RecommendedActionSessionsOperationStatus.
func (r RecommendedActionSessionsOperationStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", r.Name)
	populateTimeRFC3339(objectMap, "startTime", r.StartTime)
	populate(objectMap, "status", r.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RecommendedActionSessionsOperationStatus.
func (r *RecommendedActionSessionsOperationStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, &r.Name)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &r.StartTime)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &r.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SecurityAlertPolicyProperties.
func (s SecurityAlertPolicyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "disabledAlerts", s.DisabledAlerts)
	populate(objectMap, "emailAccountAdmins", s.EmailAccountAdmins)
	populate(objectMap, "emailAddresses", s.EmailAddresses)
	populate(objectMap, "retentionDays", s.RetentionDays)
	populate(objectMap, "state", s.State)
	populate(objectMap, "storageAccountAccessKey", s.StorageAccountAccessKey)
	populate(objectMap, "storageEndpoint", s.StorageEndpoint)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Server.
func (s Server) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServerAdministratorResourceListResult.
func (s ServerAdministratorResourceListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServerForCreate.
func (s ServerForCreate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerForCreate.
func (s *ServerForCreate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "identity":
			err = unpopulate(val, &s.Identity)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, &s.Location)
			delete(rawMsg, key)
		case "properties":
			s.Properties, err = unmarshalServerPropertiesForCreateClassification(val)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, &s.SKU)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, &s.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServerKeyListResult.
func (s ServerKeyListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServerKeyProperties.
func (s ServerKeyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "creationDate", s.CreationDate)
	populate(objectMap, "serverKeyType", s.ServerKeyType)
	populate(objectMap, "uri", s.URI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerKeyProperties.
func (s *ServerKeyProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "creationDate":
			err = unpopulateTimeRFC3339(val, &s.CreationDate)
			delete(rawMsg, key)
		case "serverKeyType":
			err = unpopulate(val, &s.ServerKeyType)
			delete(rawMsg, key)
		case "uri":
			err = unpopulate(val, &s.URI)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServerListResult.
func (s ServerListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServerProperties.
func (s ServerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "administratorLogin", s.AdministratorLogin)
	populate(objectMap, "byokEnforcement", s.ByokEnforcement)
	populateTimeRFC3339(objectMap, "earliestRestoreDate", s.EarliestRestoreDate)
	populate(objectMap, "fullyQualifiedDomainName", s.FullyQualifiedDomainName)
	populate(objectMap, "infrastructureEncryption", s.InfrastructureEncryption)
	populate(objectMap, "masterServerId", s.MasterServerID)
	populate(objectMap, "minimalTlsVersion", s.MinimalTLSVersion)
	populate(objectMap, "privateEndpointConnections", s.PrivateEndpointConnections)
	populate(objectMap, "publicNetworkAccess", s.PublicNetworkAccess)
	populate(objectMap, "replicaCapacity", s.ReplicaCapacity)
	populate(objectMap, "replicationRole", s.ReplicationRole)
	populate(objectMap, "sslEnforcement", s.SSLEnforcement)
	populate(objectMap, "storageProfile", s.StorageProfile)
	populate(objectMap, "userVisibleState", s.UserVisibleState)
	populate(objectMap, "version", s.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerProperties.
func (s *ServerProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "administratorLogin":
			err = unpopulate(val, &s.AdministratorLogin)
			delete(rawMsg, key)
		case "byokEnforcement":
			err = unpopulate(val, &s.ByokEnforcement)
			delete(rawMsg, key)
		case "earliestRestoreDate":
			err = unpopulateTimeRFC3339(val, &s.EarliestRestoreDate)
			delete(rawMsg, key)
		case "fullyQualifiedDomainName":
			err = unpopulate(val, &s.FullyQualifiedDomainName)
			delete(rawMsg, key)
		case "infrastructureEncryption":
			err = unpopulate(val, &s.InfrastructureEncryption)
			delete(rawMsg, key)
		case "masterServerId":
			err = unpopulate(val, &s.MasterServerID)
			delete(rawMsg, key)
		case "minimalTlsVersion":
			err = unpopulate(val, &s.MinimalTLSVersion)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, &s.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &s.PublicNetworkAccess)
			delete(rawMsg, key)
		case "replicaCapacity":
			err = unpopulate(val, &s.ReplicaCapacity)
			delete(rawMsg, key)
		case "replicationRole":
			err = unpopulate(val, &s.ReplicationRole)
			delete(rawMsg, key)
		case "sslEnforcement":
			err = unpopulate(val, &s.SSLEnforcement)
			delete(rawMsg, key)
		case "storageProfile":
			err = unpopulate(val, &s.StorageProfile)
			delete(rawMsg, key)
		case "userVisibleState":
			err = unpopulate(val, &s.UserVisibleState)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &s.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetServerPropertiesForCreate implements the ServerPropertiesForCreateClassification interface for type ServerPropertiesForCreate.
func (s *ServerPropertiesForCreate) GetServerPropertiesForCreate() *ServerPropertiesForCreate {
	return s
}

// GetServerPropertiesForCreate implements the ServerPropertiesForCreateClassification interface for type ServerPropertiesForDefaultCreate.
func (s *ServerPropertiesForDefaultCreate) GetServerPropertiesForCreate() *ServerPropertiesForCreate {
	return &ServerPropertiesForCreate{
		Version:                  s.Version,
		SSLEnforcement:           s.SSLEnforcement,
		MinimalTLSVersion:        s.MinimalTLSVersion,
		InfrastructureEncryption: s.InfrastructureEncryption,
		PublicNetworkAccess:      s.PublicNetworkAccess,
		StorageProfile:           s.StorageProfile,
		CreateMode:               s.CreateMode,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ServerPropertiesForDefaultCreate.
func (s ServerPropertiesForDefaultCreate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "administratorLogin", s.AdministratorLogin)
	populate(objectMap, "administratorLoginPassword", s.AdministratorLoginPassword)
	objectMap["createMode"] = CreateModeDefault
	populate(objectMap, "infrastructureEncryption", s.InfrastructureEncryption)
	populate(objectMap, "minimalTlsVersion", s.MinimalTLSVersion)
	populate(objectMap, "publicNetworkAccess", s.PublicNetworkAccess)
	populate(objectMap, "sslEnforcement", s.SSLEnforcement)
	populate(objectMap, "storageProfile", s.StorageProfile)
	populate(objectMap, "version", s.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerPropertiesForDefaultCreate.
func (s *ServerPropertiesForDefaultCreate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "administratorLogin":
			err = unpopulate(val, &s.AdministratorLogin)
			delete(rawMsg, key)
		case "administratorLoginPassword":
			err = unpopulate(val, &s.AdministratorLoginPassword)
			delete(rawMsg, key)
		case "createMode":
			err = unpopulate(val, &s.CreateMode)
			delete(rawMsg, key)
		case "infrastructureEncryption":
			err = unpopulate(val, &s.InfrastructureEncryption)
			delete(rawMsg, key)
		case "minimalTlsVersion":
			err = unpopulate(val, &s.MinimalTLSVersion)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &s.PublicNetworkAccess)
			delete(rawMsg, key)
		case "sslEnforcement":
			err = unpopulate(val, &s.SSLEnforcement)
			delete(rawMsg, key)
		case "storageProfile":
			err = unpopulate(val, &s.StorageProfile)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &s.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetServerPropertiesForCreate implements the ServerPropertiesForCreateClassification interface for type ServerPropertiesForGeoRestore.
func (s *ServerPropertiesForGeoRestore) GetServerPropertiesForCreate() *ServerPropertiesForCreate {
	return &ServerPropertiesForCreate{
		Version:                  s.Version,
		SSLEnforcement:           s.SSLEnforcement,
		MinimalTLSVersion:        s.MinimalTLSVersion,
		InfrastructureEncryption: s.InfrastructureEncryption,
		PublicNetworkAccess:      s.PublicNetworkAccess,
		StorageProfile:           s.StorageProfile,
		CreateMode:               s.CreateMode,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ServerPropertiesForGeoRestore.
func (s ServerPropertiesForGeoRestore) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["createMode"] = CreateModeGeoRestore
	populate(objectMap, "infrastructureEncryption", s.InfrastructureEncryption)
	populate(objectMap, "minimalTlsVersion", s.MinimalTLSVersion)
	populate(objectMap, "publicNetworkAccess", s.PublicNetworkAccess)
	populate(objectMap, "sslEnforcement", s.SSLEnforcement)
	populate(objectMap, "sourceServerId", s.SourceServerID)
	populate(objectMap, "storageProfile", s.StorageProfile)
	populate(objectMap, "version", s.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerPropertiesForGeoRestore.
func (s *ServerPropertiesForGeoRestore) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createMode":
			err = unpopulate(val, &s.CreateMode)
			delete(rawMsg, key)
		case "infrastructureEncryption":
			err = unpopulate(val, &s.InfrastructureEncryption)
			delete(rawMsg, key)
		case "minimalTlsVersion":
			err = unpopulate(val, &s.MinimalTLSVersion)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &s.PublicNetworkAccess)
			delete(rawMsg, key)
		case "sslEnforcement":
			err = unpopulate(val, &s.SSLEnforcement)
			delete(rawMsg, key)
		case "sourceServerId":
			err = unpopulate(val, &s.SourceServerID)
			delete(rawMsg, key)
		case "storageProfile":
			err = unpopulate(val, &s.StorageProfile)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &s.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetServerPropertiesForCreate implements the ServerPropertiesForCreateClassification interface for type ServerPropertiesForReplica.
func (s *ServerPropertiesForReplica) GetServerPropertiesForCreate() *ServerPropertiesForCreate {
	return &ServerPropertiesForCreate{
		Version:                  s.Version,
		SSLEnforcement:           s.SSLEnforcement,
		MinimalTLSVersion:        s.MinimalTLSVersion,
		InfrastructureEncryption: s.InfrastructureEncryption,
		PublicNetworkAccess:      s.PublicNetworkAccess,
		StorageProfile:           s.StorageProfile,
		CreateMode:               s.CreateMode,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ServerPropertiesForReplica.
func (s ServerPropertiesForReplica) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["createMode"] = CreateModeReplica
	populate(objectMap, "infrastructureEncryption", s.InfrastructureEncryption)
	populate(objectMap, "minimalTlsVersion", s.MinimalTLSVersion)
	populate(objectMap, "publicNetworkAccess", s.PublicNetworkAccess)
	populate(objectMap, "sslEnforcement", s.SSLEnforcement)
	populate(objectMap, "sourceServerId", s.SourceServerID)
	populate(objectMap, "storageProfile", s.StorageProfile)
	populate(objectMap, "version", s.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerPropertiesForReplica.
func (s *ServerPropertiesForReplica) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createMode":
			err = unpopulate(val, &s.CreateMode)
			delete(rawMsg, key)
		case "infrastructureEncryption":
			err = unpopulate(val, &s.InfrastructureEncryption)
			delete(rawMsg, key)
		case "minimalTlsVersion":
			err = unpopulate(val, &s.MinimalTLSVersion)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &s.PublicNetworkAccess)
			delete(rawMsg, key)
		case "sslEnforcement":
			err = unpopulate(val, &s.SSLEnforcement)
			delete(rawMsg, key)
		case "sourceServerId":
			err = unpopulate(val, &s.SourceServerID)
			delete(rawMsg, key)
		case "storageProfile":
			err = unpopulate(val, &s.StorageProfile)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &s.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetServerPropertiesForCreate implements the ServerPropertiesForCreateClassification interface for type ServerPropertiesForRestore.
func (s *ServerPropertiesForRestore) GetServerPropertiesForCreate() *ServerPropertiesForCreate {
	return &ServerPropertiesForCreate{
		Version:                  s.Version,
		SSLEnforcement:           s.SSLEnforcement,
		MinimalTLSVersion:        s.MinimalTLSVersion,
		InfrastructureEncryption: s.InfrastructureEncryption,
		PublicNetworkAccess:      s.PublicNetworkAccess,
		StorageProfile:           s.StorageProfile,
		CreateMode:               s.CreateMode,
	}
}

// MarshalJSON implements the json.Marshaller interface for type ServerPropertiesForRestore.
func (s ServerPropertiesForRestore) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["createMode"] = CreateModePointInTimeRestore
	populate(objectMap, "infrastructureEncryption", s.InfrastructureEncryption)
	populate(objectMap, "minimalTlsVersion", s.MinimalTLSVersion)
	populate(objectMap, "publicNetworkAccess", s.PublicNetworkAccess)
	populateTimeRFC3339(objectMap, "restorePointInTime", s.RestorePointInTime)
	populate(objectMap, "sslEnforcement", s.SSLEnforcement)
	populate(objectMap, "sourceServerId", s.SourceServerID)
	populate(objectMap, "storageProfile", s.StorageProfile)
	populate(objectMap, "version", s.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerPropertiesForRestore.
func (s *ServerPropertiesForRestore) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createMode":
			err = unpopulate(val, &s.CreateMode)
			delete(rawMsg, key)
		case "infrastructureEncryption":
			err = unpopulate(val, &s.InfrastructureEncryption)
			delete(rawMsg, key)
		case "minimalTlsVersion":
			err = unpopulate(val, &s.MinimalTLSVersion)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, &s.PublicNetworkAccess)
			delete(rawMsg, key)
		case "restorePointInTime":
			err = unpopulateTimeRFC3339(val, &s.RestorePointInTime)
			delete(rawMsg, key)
		case "sslEnforcement":
			err = unpopulate(val, &s.SSLEnforcement)
			delete(rawMsg, key)
		case "sourceServerId":
			err = unpopulate(val, &s.SourceServerID)
			delete(rawMsg, key)
		case "storageProfile":
			err = unpopulate(val, &s.StorageProfile)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, &s.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServerSecurityAlertPolicyListResult.
func (s ServerSecurityAlertPolicyListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServerUpdateParameters.
func (s ServerUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TagsObject.
func (t TagsObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TopQueryStatisticsInputProperties.
func (t TopQueryStatisticsInputProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregationFunction", t.AggregationFunction)
	populate(objectMap, "aggregationWindow", t.AggregationWindow)
	populate(objectMap, "numberOfTopQueries", t.NumberOfTopQueries)
	populateTimeRFC3339(objectMap, "observationEndTime", t.ObservationEndTime)
	populateTimeRFC3339(objectMap, "observationStartTime", t.ObservationStartTime)
	populate(objectMap, "observedMetric", t.ObservedMetric)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TopQueryStatisticsInputProperties.
func (t *TopQueryStatisticsInputProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aggregationFunction":
			err = unpopulate(val, &t.AggregationFunction)
			delete(rawMsg, key)
		case "aggregationWindow":
			err = unpopulate(val, &t.AggregationWindow)
			delete(rawMsg, key)
		case "numberOfTopQueries":
			err = unpopulate(val, &t.NumberOfTopQueries)
			delete(rawMsg, key)
		case "observationEndTime":
			err = unpopulateTimeRFC3339(val, &t.ObservationEndTime)
			delete(rawMsg, key)
		case "observationStartTime":
			err = unpopulateTimeRFC3339(val, &t.ObservationStartTime)
			delete(rawMsg, key)
		case "observedMetric":
			err = unpopulate(val, &t.ObservedMetric)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TopQueryStatisticsResultList.
func (t TopQueryStatisticsResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", t.NextLink)
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkRuleListResult.
func (v VirtualNetworkRuleListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WaitStatisticProperties.
func (w WaitStatisticProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", w.Count)
	populate(objectMap, "databaseName", w.DatabaseName)
	populateTimeRFC3339(objectMap, "endTime", w.EndTime)
	populate(objectMap, "eventName", w.EventName)
	populate(objectMap, "eventTypeName", w.EventTypeName)
	populate(objectMap, "queryId", w.QueryID)
	populateTimeRFC3339(objectMap, "startTime", w.StartTime)
	populate(objectMap, "totalTimeInMs", w.TotalTimeInMs)
	populate(objectMap, "userId", w.UserID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WaitStatisticProperties.
func (w *WaitStatisticProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, &w.Count)
			delete(rawMsg, key)
		case "databaseName":
			err = unpopulate(val, &w.DatabaseName)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, &w.EndTime)
			delete(rawMsg, key)
		case "eventName":
			err = unpopulate(val, &w.EventName)
			delete(rawMsg, key)
		case "eventTypeName":
			err = unpopulate(val, &w.EventTypeName)
			delete(rawMsg, key)
		case "queryId":
			err = unpopulate(val, &w.QueryID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, &w.StartTime)
			delete(rawMsg, key)
		case "totalTimeInMs":
			err = unpopulate(val, &w.TotalTimeInMs)
			delete(rawMsg, key)
		case "userId":
			err = unpopulate(val, &w.UserID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WaitStatisticsInputProperties.
func (w WaitStatisticsInputProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregationWindow", w.AggregationWindow)
	populateTimeRFC3339(objectMap, "observationEndTime", w.ObservationEndTime)
	populateTimeRFC3339(objectMap, "observationStartTime", w.ObservationStartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type WaitStatisticsInputProperties.
func (w *WaitStatisticsInputProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aggregationWindow":
			err = unpopulate(val, &w.AggregationWindow)
			delete(rawMsg, key)
		case "observationEndTime":
			err = unpopulateTimeRFC3339(val, &w.ObservationEndTime)
			delete(rawMsg, key)
		case "observationStartTime":
			err = unpopulateTimeRFC3339(val, &w.ObservationStartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type WaitStatisticsResultList.
func (w WaitStatisticsResultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", w.NextLink)
	populate(objectMap, "value", w.Value)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
