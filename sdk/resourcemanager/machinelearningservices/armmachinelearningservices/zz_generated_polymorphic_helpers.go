//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices

import "encoding/json"

func unmarshalComputeClassification(rawMsg json.RawMessage) (ComputeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ComputeClassification
	switch m["computeType"] {
	case string(ComputeTypeAKS):
		b = &AKS{}
	case string(ComputeTypeAmlCompute):
		b = &AmlCompute{}
	case string(ComputeTypeComputeInstance):
		b = &ComputeInstance{}
	case string(ComputeTypeDataFactory):
		b = &DataFactory{}
	case string(ComputeTypeDataLakeAnalytics):
		b = &DataLakeAnalytics{}
	case string(ComputeTypeDatabricks):
		b = &Databricks{}
	case string(ComputeTypeHDInsight):
		b = &HDInsight{}
	case string(ComputeTypeKubernetes):
		b = &Kubernetes{}
	case string(ComputeTypeSynapseSpark):
		b = &SynapseSpark{}
	case string(ComputeTypeVirtualMachine):
		b = &VirtualMachine{}
	default:
		b = &Compute{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalComputeSecretsClassification(rawMsg json.RawMessage) (ComputeSecretsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ComputeSecretsClassification
	switch m["computeType"] {
	case string(ComputeTypeAKS):
		b = &AksComputeSecrets{}
	case string(ComputeTypeDatabricks):
		b = &DatabricksComputeSecrets{}
	case string(ComputeTypeVirtualMachine):
		b = &VirtualMachineSecrets{}
	default:
		b = &ComputeSecrets{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
