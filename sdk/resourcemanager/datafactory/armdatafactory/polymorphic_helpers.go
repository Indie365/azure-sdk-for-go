//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import "encoding/json"

func unmarshalActivityClassification(rawMsg json.RawMessage) (ActivityClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ActivityClassification
	switch m["type"] {
	case "AppendVariable":
		b = &AppendVariableActivity{}
	case "AzureDataExplorerCommand":
		b = &AzureDataExplorerCommandActivity{}
	case "AzureFunctionActivity":
		b = &AzureFunctionActivity{}
	case "AzureMLBatchExecution":
		b = &AzureMLBatchExecutionActivity{}
	case "AzureMLExecutePipeline":
		b = &AzureMLExecutePipelineActivity{}
	case "AzureMLUpdateResource":
		b = &AzureMLUpdateResourceActivity{}
	case "Container":
		b = &ControlActivity{}
	case "Copy":
		b = &CopyActivity{}
	case "Custom":
		b = &CustomActivity{}
	case "DataLakeAnalyticsU-SQL":
		b = &DataLakeAnalyticsUSQLActivity{}
	case "DatabricksNotebook":
		b = &DatabricksNotebookActivity{}
	case "DatabricksSparkJar":
		b = &DatabricksSparkJarActivity{}
	case "DatabricksSparkPython":
		b = &DatabricksSparkPythonActivity{}
	case "Delete":
		b = &DeleteActivity{}
	case "ExecuteDataFlow":
		b = &ExecuteDataFlowActivity{}
	case "ExecutePipeline":
		b = &ExecutePipelineActivity{}
	case "ExecuteSSISPackage":
		b = &ExecuteSSISPackageActivity{}
	case "ExecuteWranglingDataflow":
		b = &ExecuteWranglingDataflowActivity{}
	case "Execution":
		b = &ExecutionActivity{}
	case "Fail":
		b = &FailActivity{}
	case "Filter":
		b = &FilterActivity{}
	case "ForEach":
		b = &ForEachActivity{}
	case "GetMetadata":
		b = &GetMetadataActivity{}
	case "HDInsightHive":
		b = &HDInsightHiveActivity{}
	case "HDInsightMapReduce":
		b = &HDInsightMapReduceActivity{}
	case "HDInsightPig":
		b = &HDInsightPigActivity{}
	case "HDInsightSpark":
		b = &HDInsightSparkActivity{}
	case "HDInsightStreaming":
		b = &HDInsightStreamingActivity{}
	case "IfCondition":
		b = &IfConditionActivity{}
	case "Lookup":
		b = &LookupActivity{}
	case "Script":
		b = &ScriptActivity{}
	case "SetVariable":
		b = &SetVariableActivity{}
	case "SparkJob":
		b = &SynapseSparkJobDefinitionActivity{}
	case "SqlServerStoredProcedure":
		b = &SQLServerStoredProcedureActivity{}
	case "Switch":
		b = &SwitchActivity{}
	case "SynapseNotebook":
		b = &SynapseNotebookActivity{}
	case "Until":
		b = &UntilActivity{}
	case "Validation":
		b = &ValidationActivity{}
	case "Wait":
		b = &WaitActivity{}
	case "WebActivity":
		b = &WebActivity{}
	case "WebHook":
		b = &WebHookActivity{}
	default:
		b = &Activity{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalActivityClassificationArray(rawMsg json.RawMessage) ([]ActivityClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ActivityClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalActivityClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalCompressionReadSettingsClassification(rawMsg json.RawMessage) (CompressionReadSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CompressionReadSettingsClassification
	switch m["type"] {
	case "TarGZipReadSettings":
		b = &TarGZipReadSettings{}
	case "TarReadSettings":
		b = &TarReadSettings{}
	case "ZipDeflateReadSettings":
		b = &ZipDeflateReadSettings{}
	default:
		b = &CompressionReadSettings{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCopySinkClassification(rawMsg json.RawMessage) (CopySinkClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CopySinkClassification
	switch m["type"] {
	case "AvroSink":
		b = &AvroSink{}
	case "AzureBlobFSSink":
		b = &AzureBlobFSSink{}
	case "AzureDataExplorerSink":
		b = &AzureDataExplorerSink{}
	case "AzureDataLakeStoreSink":
		b = &AzureDataLakeStoreSink{}
	case "AzureDatabricksDeltaLakeSink":
		b = &AzureDatabricksDeltaLakeSink{}
	case "AzureMySqlSink":
		b = &AzureMySQLSink{}
	case "AzurePostgreSqlSink":
		b = &AzurePostgreSQLSink{}
	case "AzureQueueSink":
		b = &AzureQueueSink{}
	case "AzureSearchIndexSink":
		b = &AzureSearchIndexSink{}
	case "AzureSqlSink":
		b = &AzureSQLSink{}
	case "AzureTableSink":
		b = &AzureTableSink{}
	case "BinarySink":
		b = &BinarySink{}
	case "BlobSink":
		b = &BlobSink{}
	case "CommonDataServiceForAppsSink":
		b = &CommonDataServiceForAppsSink{}
	case "CosmosDbMongoDbApiSink":
		b = &CosmosDbMongoDbAPISink{}
	case "CosmosDbSqlApiSink":
		b = &CosmosDbSQLAPISink{}
	case "DelimitedTextSink":
		b = &DelimitedTextSink{}
	case "DocumentDbCollectionSink":
		b = &DocumentDbCollectionSink{}
	case "DynamicsCrmSink":
		b = &DynamicsCrmSink{}
	case "DynamicsSink":
		b = &DynamicsSink{}
	case "FileSystemSink":
		b = &FileSystemSink{}
	case "InformixSink":
		b = &InformixSink{}
	case "JsonSink":
		b = &JSONSink{}
	case "MicrosoftAccessSink":
		b = &MicrosoftAccessSink{}
	case "MongoDbAtlasSink":
		b = &MongoDbAtlasSink{}
	case "MongoDbV2Sink":
		b = &MongoDbV2Sink{}
	case "OdbcSink":
		b = &OdbcSink{}
	case "OracleSink":
		b = &OracleSink{}
	case "OrcSink":
		b = &OrcSink{}
	case "ParquetSink":
		b = &ParquetSink{}
	case "RestSink":
		b = &RestSink{}
	case "SalesforceServiceCloudSink":
		b = &SalesforceServiceCloudSink{}
	case "SalesforceSink":
		b = &SalesforceSink{}
	case "SapCloudForCustomerSink":
		b = &SapCloudForCustomerSink{}
	case "SnowflakeSink":
		b = &SnowflakeSink{}
	case "SqlDWSink":
		b = &SQLDWSink{}
	case "SqlMISink":
		b = &SQLMISink{}
	case "SqlServerSink":
		b = &SQLServerSink{}
	case "SqlSink":
		b = &SQLSink{}
	default:
		b = &CopySink{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCopySourceClassification(rawMsg json.RawMessage) (CopySourceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CopySourceClassification
	switch m["type"] {
	case "AmazonMWSSource":
		b = &AmazonMWSSource{}
	case "AmazonRdsForOracleSource":
		b = &AmazonRdsForOracleSource{}
	case "AmazonRdsForSqlServerSource":
		b = &AmazonRdsForSQLServerSource{}
	case "AmazonRedshiftSource":
		b = &AmazonRedshiftSource{}
	case "AvroSource":
		b = &AvroSource{}
	case "AzureBlobFSSource":
		b = &AzureBlobFSSource{}
	case "AzureDataExplorerSource":
		b = &AzureDataExplorerSource{}
	case "AzureDataLakeStoreSource":
		b = &AzureDataLakeStoreSource{}
	case "AzureDatabricksDeltaLakeSource":
		b = &AzureDatabricksDeltaLakeSource{}
	case "AzureMariaDBSource":
		b = &AzureMariaDBSource{}
	case "AzureMySqlSource":
		b = &AzureMySQLSource{}
	case "AzurePostgreSqlSource":
		b = &AzurePostgreSQLSource{}
	case "AzureSqlSource":
		b = &AzureSQLSource{}
	case "AzureTableSource":
		b = &AzureTableSource{}
	case "BinarySource":
		b = &BinarySource{}
	case "BlobSource":
		b = &BlobSource{}
	case "CassandraSource":
		b = &CassandraSource{}
	case "CommonDataServiceForAppsSource":
		b = &CommonDataServiceForAppsSource{}
	case "ConcurSource":
		b = &ConcurSource{}
	case "CosmosDbMongoDbApiSource":
		b = &CosmosDbMongoDbAPISource{}
	case "CosmosDbSqlApiSource":
		b = &CosmosDbSQLAPISource{}
	case "CouchbaseSource":
		b = &CouchbaseSource{}
	case "Db2Source":
		b = &Db2Source{}
	case "DelimitedTextSource":
		b = &DelimitedTextSource{}
	case "DocumentDbCollectionSource":
		b = &DocumentDbCollectionSource{}
	case "DrillSource":
		b = &DrillSource{}
	case "DynamicsAXSource":
		b = &DynamicsAXSource{}
	case "DynamicsCrmSource":
		b = &DynamicsCrmSource{}
	case "DynamicsSource":
		b = &DynamicsSource{}
	case "EloquaSource":
		b = &EloquaSource{}
	case "ExcelSource":
		b = &ExcelSource{}
	case "FileSystemSource":
		b = &FileSystemSource{}
	case "GoogleAdWordsSource":
		b = &GoogleAdWordsSource{}
	case "GoogleBigQuerySource":
		b = &GoogleBigQuerySource{}
	case "GreenplumSource":
		b = &GreenplumSource{}
	case "HBaseSource":
		b = &HBaseSource{}
	case "HdfsSource":
		b = &HdfsSource{}
	case "HiveSource":
		b = &HiveSource{}
	case "HttpSource":
		b = &HTTPSource{}
	case "HubspotSource":
		b = &HubspotSource{}
	case "ImpalaSource":
		b = &ImpalaSource{}
	case "InformixSource":
		b = &InformixSource{}
	case "JiraSource":
		b = &JiraSource{}
	case "JsonSource":
		b = &JSONSource{}
	case "MagentoSource":
		b = &MagentoSource{}
	case "MariaDBSource":
		b = &MariaDBSource{}
	case "MarketoSource":
		b = &MarketoSource{}
	case "MicrosoftAccessSource":
		b = &MicrosoftAccessSource{}
	case "MongoDbAtlasSource":
		b = &MongoDbAtlasSource{}
	case "MongoDbSource":
		b = &MongoDbSource{}
	case "MongoDbV2Source":
		b = &MongoDbV2Source{}
	case "MySqlSource":
		b = &MySQLSource{}
	case "NetezzaSource":
		b = &NetezzaSource{}
	case "ODataSource":
		b = &ODataSource{}
	case "OdbcSource":
		b = &OdbcSource{}
	case "Office365Source":
		b = &Office365Source{}
	case "OracleServiceCloudSource":
		b = &OracleServiceCloudSource{}
	case "OracleSource":
		b = &OracleSource{}
	case "OrcSource":
		b = &OrcSource{}
	case "ParquetSource":
		b = &ParquetSource{}
	case "PaypalSource":
		b = &PaypalSource{}
	case "PhoenixSource":
		b = &PhoenixSource{}
	case "PostgreSqlSource":
		b = &PostgreSQLSource{}
	case "PrestoSource":
		b = &PrestoSource{}
	case "QuickBooksSource":
		b = &QuickBooksSource{}
	case "RelationalSource":
		b = &RelationalSource{}
	case "ResponsysSource":
		b = &ResponsysSource{}
	case "RestSource":
		b = &RestSource{}
	case "SalesforceMarketingCloudSource":
		b = &SalesforceMarketingCloudSource{}
	case "SalesforceServiceCloudSource":
		b = &SalesforceServiceCloudSource{}
	case "SalesforceSource":
		b = &SalesforceSource{}
	case "SapBwSource":
		b = &SapBwSource{}
	case "SapCloudForCustomerSource":
		b = &SapCloudForCustomerSource{}
	case "SapEccSource":
		b = &SapEccSource{}
	case "SapHanaSource":
		b = &SapHanaSource{}
	case "SapOdpSource":
		b = &SapOdpSource{}
	case "SapOpenHubSource":
		b = &SapOpenHubSource{}
	case "SapTableSource":
		b = &SapTableSource{}
	case "ServiceNowSource":
		b = &ServiceNowSource{}
	case "SharePointOnlineListSource":
		b = &SharePointOnlineListSource{}
	case "ShopifySource":
		b = &ShopifySource{}
	case "SnowflakeSource":
		b = &SnowflakeSource{}
	case "SparkSource":
		b = &SparkSource{}
	case "SqlDWSource":
		b = &SQLDWSource{}
	case "SqlMISource":
		b = &SQLMISource{}
	case "SqlServerSource":
		b = &SQLServerSource{}
	case "SqlSource":
		b = &SQLSource{}
	case "SquareSource":
		b = &SquareSource{}
	case "SybaseSource":
		b = &SybaseSource{}
	case "TabularSource":
		b = &TabularSource{}
	case "TeradataSource":
		b = &TeradataSource{}
	case "VerticaSource":
		b = &VerticaSource{}
	case "WebSource":
		b = &WebSource{}
	case "XeroSource":
		b = &XeroSource{}
	case "XmlSource":
		b = &XMLSource{}
	case "ZohoSource":
		b = &ZohoSource{}
	default:
		b = &CopySource{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCredentialClassification(rawMsg json.RawMessage) (CredentialClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CredentialClassification
	switch m["type"] {
	case "ManagedIdentity":
		b = &ManagedIdentityCredential{}
	case "ServicePrincipal":
		b = &ServicePrincipalCredential{}
	default:
		b = &Credential{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCustomSetupBaseClassification(rawMsg json.RawMessage) (CustomSetupBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CustomSetupBaseClassification
	switch m["type"] {
	case "AzPowerShellSetup":
		b = &AzPowerShellSetup{}
	case "CmdkeySetup":
		b = &CmdkeySetup{}
	case "ComponentSetup":
		b = &ComponentSetup{}
	case "EnvironmentVariableSetup":
		b = &EnvironmentVariableSetup{}
	default:
		b = &CustomSetupBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCustomSetupBaseClassificationArray(rawMsg json.RawMessage) ([]CustomSetupBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]CustomSetupBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalCustomSetupBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDataFlowClassification(rawMsg json.RawMessage) (DataFlowClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataFlowClassification
	switch m["type"] {
	case "Flowlet":
		b = &Flowlet{}
	case "MappingDataFlow":
		b = &MappingDataFlow{}
	case "WranglingDataFlow":
		b = &WranglingDataFlow{}
	default:
		b = &DataFlow{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatasetClassification(rawMsg json.RawMessage) (DatasetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatasetClassification
	switch m["type"] {
	case "AmazonMWSObject":
		b = &AmazonMWSObjectDataset{}
	case "AmazonRdsForOracleTable":
		b = &AmazonRdsForOracleTableDataset{}
	case "AmazonRdsForSqlServerTable":
		b = &AmazonRdsForSQLServerTableDataset{}
	case "AmazonRedshiftTable":
		b = &AmazonRedshiftTableDataset{}
	case "AmazonS3Object":
		b = &AmazonS3Dataset{}
	case "Avro":
		b = &AvroDataset{}
	case "AzureBlob":
		b = &AzureBlobDataset{}
	case "AzureBlobFSFile":
		b = &AzureBlobFSDataset{}
	case "AzureDataExplorerTable":
		b = &AzureDataExplorerTableDataset{}
	case "AzureDataLakeStoreFile":
		b = &AzureDataLakeStoreDataset{}
	case "AzureDatabricksDeltaLakeDataset":
		b = &AzureDatabricksDeltaLakeDataset{}
	case "AzureMariaDBTable":
		b = &AzureMariaDBTableDataset{}
	case "AzureMySqlTable":
		b = &AzureMySQLTableDataset{}
	case "AzurePostgreSqlTable":
		b = &AzurePostgreSQLTableDataset{}
	case "AzureSearchIndex":
		b = &AzureSearchIndexDataset{}
	case "AzureSqlDWTable":
		b = &AzureSQLDWTableDataset{}
	case "AzureSqlMITable":
		b = &AzureSQLMITableDataset{}
	case "AzureSqlTable":
		b = &AzureSQLTableDataset{}
	case "AzureTable":
		b = &AzureTableDataset{}
	case "Binary":
		b = &BinaryDataset{}
	case "CassandraTable":
		b = &CassandraTableDataset{}
	case "CommonDataServiceForAppsEntity":
		b = &CommonDataServiceForAppsEntityDataset{}
	case "ConcurObject":
		b = &ConcurObjectDataset{}
	case "CosmosDbMongoDbApiCollection":
		b = &CosmosDbMongoDbAPICollectionDataset{}
	case "CosmosDbSqlApiCollection":
		b = &CosmosDbSQLAPICollectionDataset{}
	case "CouchbaseTable":
		b = &CouchbaseTableDataset{}
	case "CustomDataset":
		b = &CustomDataset{}
	case "Db2Table":
		b = &Db2TableDataset{}
	case "DelimitedText":
		b = &DelimitedTextDataset{}
	case "DocumentDbCollection":
		b = &DocumentDbCollectionDataset{}
	case "DrillTable":
		b = &DrillTableDataset{}
	case "DynamicsAXResource":
		b = &DynamicsAXResourceDataset{}
	case "DynamicsCrmEntity":
		b = &DynamicsCrmEntityDataset{}
	case "DynamicsEntity":
		b = &DynamicsEntityDataset{}
	case "EloquaObject":
		b = &EloquaObjectDataset{}
	case "Excel":
		b = &ExcelDataset{}
	case "FileShare":
		b = &FileShareDataset{}
	case "GoogleAdWordsObject":
		b = &GoogleAdWordsObjectDataset{}
	case "GoogleBigQueryObject":
		b = &GoogleBigQueryObjectDataset{}
	case "GreenplumTable":
		b = &GreenplumTableDataset{}
	case "HBaseObject":
		b = &HBaseObjectDataset{}
	case "HiveObject":
		b = &HiveObjectDataset{}
	case "HttpFile":
		b = &HTTPDataset{}
	case "HubspotObject":
		b = &HubspotObjectDataset{}
	case "ImpalaObject":
		b = &ImpalaObjectDataset{}
	case "InformixTable":
		b = &InformixTableDataset{}
	case "JiraObject":
		b = &JiraObjectDataset{}
	case "Json":
		b = &JSONDataset{}
	case "MagentoObject":
		b = &MagentoObjectDataset{}
	case "MariaDBTable":
		b = &MariaDBTableDataset{}
	case "MarketoObject":
		b = &MarketoObjectDataset{}
	case "MicrosoftAccessTable":
		b = &MicrosoftAccessTableDataset{}
	case "MongoDbAtlasCollection":
		b = &MongoDbAtlasCollectionDataset{}
	case "MongoDbCollection":
		b = &MongoDbCollectionDataset{}
	case "MongoDbV2Collection":
		b = &MongoDbV2CollectionDataset{}
	case "MySqlTable":
		b = &MySQLTableDataset{}
	case "NetezzaTable":
		b = &NetezzaTableDataset{}
	case "ODataResource":
		b = &ODataResourceDataset{}
	case "OdbcTable":
		b = &OdbcTableDataset{}
	case "Office365Table":
		b = &Office365Dataset{}
	case "OracleServiceCloudObject":
		b = &OracleServiceCloudObjectDataset{}
	case "OracleTable":
		b = &OracleTableDataset{}
	case "Orc":
		b = &OrcDataset{}
	case "Parquet":
		b = &ParquetDataset{}
	case "PaypalObject":
		b = &PaypalObjectDataset{}
	case "PhoenixObject":
		b = &PhoenixObjectDataset{}
	case "PostgreSqlTable":
		b = &PostgreSQLTableDataset{}
	case "PrestoObject":
		b = &PrestoObjectDataset{}
	case "QuickBooksObject":
		b = &QuickBooksObjectDataset{}
	case "RelationalTable":
		b = &RelationalTableDataset{}
	case "ResponsysObject":
		b = &ResponsysObjectDataset{}
	case "RestResource":
		b = &RestResourceDataset{}
	case "SalesforceMarketingCloudObject":
		b = &SalesforceMarketingCloudObjectDataset{}
	case "SalesforceObject":
		b = &SalesforceObjectDataset{}
	case "SalesforceServiceCloudObject":
		b = &SalesforceServiceCloudObjectDataset{}
	case "SapBwCube":
		b = &SapBwCubeDataset{}
	case "SapCloudForCustomerResource":
		b = &SapCloudForCustomerResourceDataset{}
	case "SapEccResource":
		b = &SapEccResourceDataset{}
	case "SapHanaTable":
		b = &SapHanaTableDataset{}
	case "SapOdpResource":
		b = &SapOdpResourceDataset{}
	case "SapOpenHubTable":
		b = &SapOpenHubTableDataset{}
	case "SapTableResource":
		b = &SapTableResourceDataset{}
	case "ServiceNowObject":
		b = &ServiceNowObjectDataset{}
	case "SharePointOnlineListResource":
		b = &SharePointOnlineListResourceDataset{}
	case "ShopifyObject":
		b = &ShopifyObjectDataset{}
	case "SnowflakeTable":
		b = &SnowflakeDataset{}
	case "SparkObject":
		b = &SparkObjectDataset{}
	case "SqlServerTable":
		b = &SQLServerTableDataset{}
	case "SquareObject":
		b = &SquareObjectDataset{}
	case "SybaseTable":
		b = &SybaseTableDataset{}
	case "TeradataTable":
		b = &TeradataTableDataset{}
	case "VerticaTable":
		b = &VerticaTableDataset{}
	case "WebTable":
		b = &WebTableDataset{}
	case "XeroObject":
		b = &XeroObjectDataset{}
	case "Xml":
		b = &XMLDataset{}
	case "ZohoObject":
		b = &ZohoObjectDataset{}
	default:
		b = &Dataset{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatasetLocationClassification(rawMsg json.RawMessage) (DatasetLocationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatasetLocationClassification
	switch m["type"] {
	case "AmazonS3CompatibleLocation":
		b = &AmazonS3CompatibleLocation{}
	case "AmazonS3Location":
		b = &AmazonS3Location{}
	case "AzureBlobFSLocation":
		b = &AzureBlobFSLocation{}
	case "AzureBlobStorageLocation":
		b = &AzureBlobStorageLocation{}
	case "AzureDataLakeStoreLocation":
		b = &AzureDataLakeStoreLocation{}
	case "AzureFileStorageLocation":
		b = &AzureFileStorageLocation{}
	case "FileServerLocation":
		b = &FileServerLocation{}
	case "FtpServerLocation":
		b = &FtpServerLocation{}
	case "GoogleCloudStorageLocation":
		b = &GoogleCloudStorageLocation{}
	case "HdfsLocation":
		b = &HdfsLocation{}
	case "HttpServerLocation":
		b = &HTTPServerLocation{}
	case "OracleCloudStorageLocation":
		b = &OracleCloudStorageLocation{}
	case "SftpLocation":
		b = &SftpLocation{}
	default:
		b = &DatasetLocation{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatasetStorageFormatClassification(rawMsg json.RawMessage) (DatasetStorageFormatClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatasetStorageFormatClassification
	switch m["type"] {
	case "AvroFormat":
		b = &AvroFormat{}
	case "JsonFormat":
		b = &JSONFormat{}
	case "OrcFormat":
		b = &OrcFormat{}
	case "ParquetFormat":
		b = &ParquetFormat{}
	case "TextFormat":
		b = &TextFormat{}
	default:
		b = &DatasetStorageFormat{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDependencyReferenceClassification(rawMsg json.RawMessage) (DependencyReferenceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DependencyReferenceClassification
	switch m["type"] {
	case "SelfDependencyTumblingWindowTriggerReference":
		b = &SelfDependencyTumblingWindowTriggerReference{}
	case "TriggerDependencyReference":
		b = &TriggerDependencyReference{}
	case "TumblingWindowTriggerDependencyReference":
		b = &TumblingWindowTriggerDependencyReference{}
	default:
		b = &DependencyReference{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDependencyReferenceClassificationArray(rawMsg json.RawMessage) ([]DependencyReferenceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DependencyReferenceClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDependencyReferenceClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFactoryRepoConfigurationClassification(rawMsg json.RawMessage) (FactoryRepoConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FactoryRepoConfigurationClassification
	switch m["type"] {
	case "FactoryGitHubConfiguration":
		b = &FactoryGitHubConfiguration{}
	case "FactoryVSTSConfiguration":
		b = &FactoryVSTSConfiguration{}
	default:
		b = &FactoryRepoConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFormatReadSettingsClassification(rawMsg json.RawMessage) (FormatReadSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FormatReadSettingsClassification
	switch m["type"] {
	case "BinaryReadSettings":
		b = &BinaryReadSettings{}
	case "DelimitedTextReadSettings":
		b = &DelimitedTextReadSettings{}
	case "JsonReadSettings":
		b = &JSONReadSettings{}
	case "XmlReadSettings":
		b = &XMLReadSettings{}
	default:
		b = &FormatReadSettings{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalIntegrationRuntimeClassification(rawMsg json.RawMessage) (IntegrationRuntimeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b IntegrationRuntimeClassification
	switch m["type"] {
	case string(IntegrationRuntimeTypeManaged):
		b = &ManagedIntegrationRuntime{}
	case string(IntegrationRuntimeTypeSelfHosted):
		b = &SelfHostedIntegrationRuntime{}
	default:
		b = &IntegrationRuntime{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalIntegrationRuntimeStatusClassification(rawMsg json.RawMessage) (IntegrationRuntimeStatusClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b IntegrationRuntimeStatusClassification
	switch m["type"] {
	case string(IntegrationRuntimeTypeManaged):
		b = &ManagedIntegrationRuntimeStatus{}
	case string(IntegrationRuntimeTypeSelfHosted):
		b = &SelfHostedIntegrationRuntimeStatus{}
	default:
		b = &IntegrationRuntimeStatus{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalLinkedIntegrationRuntimeTypeClassification(rawMsg json.RawMessage) (LinkedIntegrationRuntimeTypeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b LinkedIntegrationRuntimeTypeClassification
	switch m["authorizationType"] {
	case "Key":
		b = &LinkedIntegrationRuntimeKeyAuthorization{}
	case "RBAC":
		b = &LinkedIntegrationRuntimeRbacAuthorization{}
	default:
		b = &LinkedIntegrationRuntimeType{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalLinkedServiceClassification(rawMsg json.RawMessage) (LinkedServiceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b LinkedServiceClassification
	switch m["type"] {
	case "AmazonMWS":
		b = &AmazonMWSLinkedService{}
	case "AmazonRdsForOracle":
		b = &AmazonRdsForOracleLinkedService{}
	case "AmazonRdsForSqlServer":
		b = &AmazonRdsForSQLServerLinkedService{}
	case "AmazonRedshift":
		b = &AmazonRedshiftLinkedService{}
	case "AmazonS3":
		b = &AmazonS3LinkedService{}
	case "AmazonS3Compatible":
		b = &AmazonS3CompatibleLinkedService{}
	case "AppFigures":
		b = &AppFiguresLinkedService{}
	case "Asana":
		b = &AsanaLinkedService{}
	case "AzureBatch":
		b = &AzureBatchLinkedService{}
	case "AzureBlobFS":
		b = &AzureBlobFSLinkedService{}
	case "AzureBlobStorage":
		b = &AzureBlobStorageLinkedService{}
	case "AzureDataExplorer":
		b = &AzureDataExplorerLinkedService{}
	case "AzureDataLakeAnalytics":
		b = &AzureDataLakeAnalyticsLinkedService{}
	case "AzureDataLakeStore":
		b = &AzureDataLakeStoreLinkedService{}
	case "AzureDatabricks":
		b = &AzureDatabricksLinkedService{}
	case "AzureDatabricksDeltaLake":
		b = &AzureDatabricksDeltaLakeLinkedService{}
	case "AzureFileStorage":
		b = &AzureFileStorageLinkedService{}
	case "AzureFunction":
		b = &AzureFunctionLinkedService{}
	case "AzureKeyVault":
		b = &AzureKeyVaultLinkedService{}
	case "AzureML":
		b = &AzureMLLinkedService{}
	case "AzureMLService":
		b = &AzureMLServiceLinkedService{}
	case "AzureMariaDB":
		b = &AzureMariaDBLinkedService{}
	case "AzureMySql":
		b = &AzureMySQLLinkedService{}
	case "AzurePostgreSql":
		b = &AzurePostgreSQLLinkedService{}
	case "AzureSearch":
		b = &AzureSearchLinkedService{}
	case "AzureSqlDW":
		b = &AzureSQLDWLinkedService{}
	case "AzureSqlDatabase":
		b = &AzureSQLDatabaseLinkedService{}
	case "AzureSqlMI":
		b = &AzureSQLMILinkedService{}
	case "AzureStorage":
		b = &AzureStorageLinkedService{}
	case "AzureSynapseArtifacts":
		b = &AzureSynapseArtifactsLinkedService{}
	case "AzureTableStorage":
		b = &AzureTableStorageLinkedService{}
	case "Cassandra":
		b = &CassandraLinkedService{}
	case "CommonDataServiceForApps":
		b = &CommonDataServiceForAppsLinkedService{}
	case "Concur":
		b = &ConcurLinkedService{}
	case "CosmosDb":
		b = &CosmosDbLinkedService{}
	case "CosmosDbMongoDbApi":
		b = &CosmosDbMongoDbAPILinkedService{}
	case "Couchbase":
		b = &CouchbaseLinkedService{}
	case "CustomDataSource":
		b = &CustomDataSourceLinkedService{}
	case "Dataworld":
		b = &DataworldLinkedService{}
	case "Db2":
		b = &Db2LinkedService{}
	case "Drill":
		b = &DrillLinkedService{}
	case "Dynamics":
		b = &DynamicsLinkedService{}
	case "DynamicsAX":
		b = &DynamicsAXLinkedService{}
	case "DynamicsCrm":
		b = &DynamicsCrmLinkedService{}
	case "Eloqua":
		b = &EloquaLinkedService{}
	case "FileServer":
		b = &FileServerLinkedService{}
	case "FtpServer":
		b = &FtpServerLinkedService{}
	case "GoogleAdWords":
		b = &GoogleAdWordsLinkedService{}
	case "GoogleBigQuery":
		b = &GoogleBigQueryLinkedService{}
	case "GoogleCloudStorage":
		b = &GoogleCloudStorageLinkedService{}
	case "GoogleSheets":
		b = &GoogleSheetsLinkedService{}
	case "Greenplum":
		b = &GreenplumLinkedService{}
	case "HBase":
		b = &HBaseLinkedService{}
	case "HDInsight":
		b = &HDInsightLinkedService{}
	case "HDInsightOnDemand":
		b = &HDInsightOnDemandLinkedService{}
	case "Hdfs":
		b = &HdfsLinkedService{}
	case "Hive":
		b = &HiveLinkedService{}
	case "HttpServer":
		b = &HTTPLinkedService{}
	case "Hubspot":
		b = &HubspotLinkedService{}
	case "Impala":
		b = &ImpalaLinkedService{}
	case "Informix":
		b = &InformixLinkedService{}
	case "Jira":
		b = &JiraLinkedService{}
	case "Magento":
		b = &MagentoLinkedService{}
	case "MariaDB":
		b = &MariaDBLinkedService{}
	case "Marketo":
		b = &MarketoLinkedService{}
	case "MicrosoftAccess":
		b = &MicrosoftAccessLinkedService{}
	case "MongoDb":
		b = &MongoDbLinkedService{}
	case "MongoDbAtlas":
		b = &MongoDbAtlasLinkedService{}
	case "MongoDbV2":
		b = &MongoDbV2LinkedService{}
	case "MySql":
		b = &MySQLLinkedService{}
	case "Netezza":
		b = &NetezzaLinkedService{}
	case "OData":
		b = &ODataLinkedService{}
	case "Odbc":
		b = &OdbcLinkedService{}
	case "Office365":
		b = &Office365LinkedService{}
	case "Oracle":
		b = &OracleLinkedService{}
	case "OracleCloudStorage":
		b = &OracleCloudStorageLinkedService{}
	case "OracleServiceCloud":
		b = &OracleServiceCloudLinkedService{}
	case "Paypal":
		b = &PaypalLinkedService{}
	case "Phoenix":
		b = &PhoenixLinkedService{}
	case "PostgreSql":
		b = &PostgreSQLLinkedService{}
	case "Presto":
		b = &PrestoLinkedService{}
	case "QuickBooks":
		b = &QuickBooksLinkedService{}
	case "Quickbase":
		b = &QuickbaseLinkedService{}
	case "Responsys":
		b = &ResponsysLinkedService{}
	case "RestService":
		b = &RestServiceLinkedService{}
	case "Salesforce":
		b = &SalesforceLinkedService{}
	case "SalesforceMarketingCloud":
		b = &SalesforceMarketingCloudLinkedService{}
	case "SalesforceServiceCloud":
		b = &SalesforceServiceCloudLinkedService{}
	case "SapBW":
		b = &SapBWLinkedService{}
	case "SapCloudForCustomer":
		b = &SapCloudForCustomerLinkedService{}
	case "SapEcc":
		b = &SapEccLinkedService{}
	case "SapHana":
		b = &SapHanaLinkedService{}
	case "SapOdp":
		b = &SapOdpLinkedService{}
	case "SapOpenHub":
		b = &SapOpenHubLinkedService{}
	case "SapTable":
		b = &SapTableLinkedService{}
	case "ServiceNow":
		b = &ServiceNowLinkedService{}
	case "Sftp":
		b = &SftpServerLinkedService{}
	case "SharePointOnlineList":
		b = &SharePointOnlineListLinkedService{}
	case "Shopify":
		b = &ShopifyLinkedService{}
	case "Smartsheet":
		b = &SmartsheetLinkedService{}
	case "Snowflake":
		b = &SnowflakeLinkedService{}
	case "Spark":
		b = &SparkLinkedService{}
	case "SqlServer":
		b = &SQLServerLinkedService{}
	case "Square":
		b = &SquareLinkedService{}
	case "Sybase":
		b = &SybaseLinkedService{}
	case "TeamDesk":
		b = &TeamDeskLinkedService{}
	case "Teradata":
		b = &TeradataLinkedService{}
	case "Twilio":
		b = &TwilioLinkedService{}
	case "Vertica":
		b = &VerticaLinkedService{}
	case "Web":
		b = &WebLinkedService{}
	case "Xero":
		b = &XeroLinkedService{}
	case "Zendesk":
		b = &ZendeskLinkedService{}
	case "Zoho":
		b = &ZohoLinkedService{}
	default:
		b = &LinkedService{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSecretBaseClassification(rawMsg json.RawMessage) (SecretBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SecretBaseClassification
	switch m["type"] {
	case "AzureKeyVaultSecret":
		b = &AzureKeyVaultSecretReference{}
	case "SecureString":
		b = &SecureString{}
	default:
		b = &SecretBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSsisObjectMetadataClassification(rawMsg json.RawMessage) (SsisObjectMetadataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SsisObjectMetadataClassification
	switch m["type"] {
	case string(SsisObjectMetadataTypeEnvironment):
		b = &SsisEnvironment{}
	case string(SsisObjectMetadataTypeFolder):
		b = &SsisFolder{}
	case string(SsisObjectMetadataTypePackage):
		b = &SsisPackage{}
	case string(SsisObjectMetadataTypeProject):
		b = &SsisProject{}
	default:
		b = &SsisObjectMetadata{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSsisObjectMetadataClassificationArray(rawMsg json.RawMessage) ([]SsisObjectMetadataClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]SsisObjectMetadataClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalSsisObjectMetadataClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalStoreReadSettingsClassification(rawMsg json.RawMessage) (StoreReadSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b StoreReadSettingsClassification
	switch m["type"] {
	case "AmazonS3CompatibleReadSettings":
		b = &AmazonS3CompatibleReadSettings{}
	case "AmazonS3ReadSettings":
		b = &AmazonS3ReadSettings{}
	case "AzureBlobFSReadSettings":
		b = &AzureBlobFSReadSettings{}
	case "AzureBlobStorageReadSettings":
		b = &AzureBlobStorageReadSettings{}
	case "AzureDataLakeStoreReadSettings":
		b = &AzureDataLakeStoreReadSettings{}
	case "AzureFileStorageReadSettings":
		b = &AzureFileStorageReadSettings{}
	case "FileServerReadSettings":
		b = &FileServerReadSettings{}
	case "FtpReadSettings":
		b = &FtpReadSettings{}
	case "GoogleCloudStorageReadSettings":
		b = &GoogleCloudStorageReadSettings{}
	case "HdfsReadSettings":
		b = &HdfsReadSettings{}
	case "HttpReadSettings":
		b = &HTTPReadSettings{}
	case "OracleCloudStorageReadSettings":
		b = &OracleCloudStorageReadSettings{}
	case "SftpReadSettings":
		b = &SftpReadSettings{}
	default:
		b = &StoreReadSettings{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalStoreWriteSettingsClassification(rawMsg json.RawMessage) (StoreWriteSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b StoreWriteSettingsClassification
	switch m["type"] {
	case "AzureBlobFSWriteSettings":
		b = &AzureBlobFSWriteSettings{}
	case "AzureBlobStorageWriteSettings":
		b = &AzureBlobStorageWriteSettings{}
	case "AzureDataLakeStoreWriteSettings":
		b = &AzureDataLakeStoreWriteSettings{}
	case "AzureFileStorageWriteSettings":
		b = &AzureFileStorageWriteSettings{}
	case "FileServerWriteSettings":
		b = &FileServerWriteSettings{}
	case "SftpWriteSettings":
		b = &SftpWriteSettings{}
	default:
		b = &StoreWriteSettings{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTriggerClassification(rawMsg json.RawMessage) (TriggerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TriggerClassification
	switch m["type"] {
	case "BlobEventsTrigger":
		b = &BlobEventsTrigger{}
	case "BlobTrigger":
		b = &BlobTrigger{}
	case "ChainingTrigger":
		b = &ChainingTrigger{}
	case "CustomEventsTrigger":
		b = &CustomEventsTrigger{}
	case "MultiplePipelineTrigger":
		b = &MultiplePipelineTrigger{}
	case "RerunTumblingWindowTrigger":
		b = &RerunTumblingWindowTrigger{}
	case "ScheduleTrigger":
		b = &ScheduleTrigger{}
	case "TumblingWindowTrigger":
		b = &TumblingWindowTrigger{}
	default:
		b = &Trigger{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalWebLinkedServiceTypePropertiesClassification(rawMsg json.RawMessage) (WebLinkedServiceTypePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b WebLinkedServiceTypePropertiesClassification
	switch m["authenticationType"] {
	case string(WebAuthenticationTypeAnonymous):
		b = &WebAnonymousAuthentication{}
	case string(WebAuthenticationTypeBasic):
		b = &WebBasicAuthentication{}
	case string(WebAuthenticationTypeClientCertificate):
		b = &WebClientCertificateAuthentication{}
	default:
		b = &WebLinkedServiceTypeProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

