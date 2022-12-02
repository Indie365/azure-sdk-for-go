//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

import "encoding/json"

func unmarshalRoleManagementPolicyRuleClassification(rawMsg json.RawMessage) (RoleManagementPolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RoleManagementPolicyRuleClassification
	switch m["ruleType"] {
	case string(RoleManagementPolicyRuleTypeRoleManagementPolicyApprovalRule):
		b = &RoleManagementPolicyApprovalRule{}
	case string(RoleManagementPolicyRuleTypeRoleManagementPolicyAuthenticationContextRule):
		b = &RoleManagementPolicyAuthenticationContextRule{}
	case string(RoleManagementPolicyRuleTypeRoleManagementPolicyEnablementRule):
		b = &RoleManagementPolicyEnablementRule{}
	case string(RoleManagementPolicyRuleTypeRoleManagementPolicyExpirationRule):
		b = &RoleManagementPolicyExpirationRule{}
	case string(RoleManagementPolicyRuleTypeRoleManagementPolicyNotificationRule):
		b = &RoleManagementPolicyNotificationRule{}
	default:
		b = &RoleManagementPolicyRule{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRoleManagementPolicyRuleClassificationArray(rawMsg json.RawMessage) ([]RoleManagementPolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RoleManagementPolicyRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRoleManagementPolicyRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}