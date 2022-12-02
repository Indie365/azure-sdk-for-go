//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

import "encoding/json"

func unmarshalImageTemplateCustomizerClassification(rawMsg json.RawMessage) (ImageTemplateCustomizerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageTemplateCustomizerClassification
	switch m["type"] {
	case "File":
		b = &ImageTemplateFileCustomizer{}
	case "PowerShell":
		b = &ImageTemplatePowerShellCustomizer{}
	case "Shell":
		b = &ImageTemplateShellCustomizer{}
	case "WindowsRestart":
		b = &ImageTemplateRestartCustomizer{}
	case "WindowsUpdate":
		b = &ImageTemplateWindowsUpdateCustomizer{}
	default:
		b = &ImageTemplateCustomizer{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalImageTemplateCustomizerClassificationArray(rawMsg json.RawMessage) ([]ImageTemplateCustomizerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ImageTemplateCustomizerClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalImageTemplateCustomizerClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalImageTemplateDistributorClassification(rawMsg json.RawMessage) (ImageTemplateDistributorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageTemplateDistributorClassification
	switch m["type"] {
	case "ManagedImage":
		b = &ImageTemplateManagedImageDistributor{}
	case "SharedImage":
		b = &ImageTemplateSharedImageDistributor{}
	case "VHD":
		b = &ImageTemplateVhdDistributor{}
	default:
		b = &ImageTemplateDistributor{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalImageTemplateDistributorClassificationArray(rawMsg json.RawMessage) ([]ImageTemplateDistributorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ImageTemplateDistributorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalImageTemplateDistributorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalImageTemplateInVMValidatorClassification(rawMsg json.RawMessage) (ImageTemplateInVMValidatorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageTemplateInVMValidatorClassification
	switch m["type"] {
	case "PowerShell":
		b = &ImageTemplatePowerShellValidator{}
	case "Shell":
		b = &ImageTemplateShellValidator{}
	default:
		b = &ImageTemplateInVMValidator{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalImageTemplateInVMValidatorClassificationArray(rawMsg json.RawMessage) ([]ImageTemplateInVMValidatorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ImageTemplateInVMValidatorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalImageTemplateInVMValidatorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalImageTemplateSourceClassification(rawMsg json.RawMessage) (ImageTemplateSourceClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageTemplateSourceClassification
	switch m["type"] {
	case "ManagedImage":
		b = &ImageTemplateManagedImageSource{}
	case "PlatformImage":
		b = &ImageTemplatePlatformImageSource{}
	case "SharedImageVersion":
		b = &ImageTemplateSharedImageVersionSource{}
	default:
		b = &ImageTemplateSource{}
	}
	return b, json.Unmarshal(rawMsg, b)
}