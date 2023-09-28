//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// Client contains the methods for the OpenAI group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal *azcore.Client
	clientData
}

// beginAzureBatchImageGeneration - Starts the generation of a batch of images from a text caption
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - options - beginAzureBatchImageGenerationOptions contains the optional parameters for the Client.beginAzureBatchImageGeneration
//     method.
func (client *Client) beginAzureBatchImageGeneration(ctx context.Context, body ImageGenerationOptions, options *beginAzureBatchImageGenerationOptions) (*runtime.Poller[azureBatchImageGenerationInternalResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.azureBatchImageGenerationInternal(ctx, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[azureBatchImageGenerationInternalResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[azureBatchImageGenerationInternalResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// AzureBatchImageGenerationInternal - Starts the generation of a batch of images from a text caption
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *Client) azureBatchImageGenerationInternal(ctx context.Context, body ImageGenerationOptions, options *beginAzureBatchImageGenerationOptions) (*http.Response, error) {
	var err error
	req, err := client.azureBatchImageGenerationInternalCreateRequest(ctx, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = client.newError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// azureBatchImageGenerationInternalCreateRequest creates the AzureBatchImageGenerationInternal request.
func (client *Client) azureBatchImageGenerationInternalCreateRequest(ctx context.Context, body ImageGenerationOptions, options *beginAzureBatchImageGenerationOptions) (*policy.Request, error) {
	urlPath := "/images/generations:submit"
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.formatURL(urlPath, getDeployment(body)))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getAudioTranscriptionInternal - Gets transcribed text and associated metadata from provided spoken audio data. Audio will
// be transcribed in the written language corresponding to the language it was spoken in. Gets transcribed text
// and associated metadata from provided spoken audio data. Audio will be transcribed in the written language corresponding
// to the language it was spoken in.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - file - The audio data to transcribe. This must be the binary content of a file in one of the supported media formats: flac,
//     mp3, mp4, mpeg, mpga, m4a, ogg, wav, webm.
//   - options - getAudioTranscriptionInternalOptions contains the optional parameters for the Client.getAudioTranscriptionInternal
//     method.
func (client *Client) getAudioTranscriptionInternal(ctx context.Context, file []byte, options *getAudioTranscriptionInternalOptions) (getAudioTranscriptionInternalResponse, error) {
	var err error
	req, err := client.getAudioTranscriptionInternalCreateRequest(ctx, file, options)
	if err != nil {
		return getAudioTranscriptionInternalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return getAudioTranscriptionInternalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = client.newError(httpResp)
		return getAudioTranscriptionInternalResponse{}, err
	}
	resp, err := getAudioTranscriptionInternalHandleResponse(httpResp)
	return resp, err
}

// getAudioTranscriptionInternalCreateRequest creates the getAudioTranscriptionInternal request.
func (client *Client) getAudioTranscriptionInternalCreateRequest(ctx context.Context, file []byte, body *getAudioTranscriptionInternalOptions) (*policy.Request, error) {
	urlPath := "audio/transcriptions"
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.formatURL(urlPath, getDeployment(body)))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := setMultipartFormData(req, file, *body); err != nil {
		return nil, err
	}
	return req, nil
}

// getAudioTranscriptionInternalHandleResponse handles the getAudioTranscriptionInternal response.
func (client *Client) getAudioTranscriptionInternalHandleResponse(resp *http.Response) (getAudioTranscriptionInternalResponse, error) {
	result := getAudioTranscriptionInternalResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AudioTranscription); err != nil {
		return getAudioTranscriptionInternalResponse{}, err
	}
	return result, nil
}

// getAudioTranslationInternal - Gets English language transcribed text and associated metadata from provided spoken audio
// data. Gets English language transcribed text and associated metadata from provided spoken audio data.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - file - The audio data to translate. This must be the binary content of a file in one of the supported media formats: flac,
//     mp3, mp4, mpeg, mpga, m4a, ogg, wav, webm.
//   - options - getAudioTranslationInternalOptions contains the optional parameters for the Client.getAudioTranslationInternal
//     method.
func (client *Client) getAudioTranslationInternal(ctx context.Context, file []byte, options *getAudioTranslationInternalOptions) (getAudioTranslationInternalResponse, error) {
	var err error
	req, err := client.getAudioTranslationInternalCreateRequest(ctx, file, options)
	if err != nil {
		return getAudioTranslationInternalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return getAudioTranslationInternalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = client.newError(httpResp)
		return getAudioTranslationInternalResponse{}, err
	}
	resp, err := getAudioTranslationInternalHandleResponse(httpResp)
	return resp, err
}

// getAudioTranslationInternalCreateRequest creates the getAudioTranslationInternal request.
func (client *Client) getAudioTranslationInternalCreateRequest(ctx context.Context, file []byte, body *getAudioTranslationInternalOptions) (*policy.Request, error) {
	urlPath := "audio/translations"
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.formatURL(urlPath, getDeployment(body)))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := setMultipartFormData(req, file, *body); err != nil {
		return nil, err
	}
	return req, nil
}

// getAudioTranslationInternalHandleResponse handles the getAudioTranslationInternal response.
func (client *Client) getAudioTranslationInternalHandleResponse(resp *http.Response) (getAudioTranslationInternalResponse, error) {
	result := getAudioTranslationInternalResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AudioTranscription); err != nil {
		return getAudioTranslationInternalResponse{}, err
	}
	return result, nil
}

// getChatCompletions - Gets chat completions for the provided chat messages. Completions support a wide variety of tasks
// and generate text that continues from or "completes" provided prompt data.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - options - GetChatCompletionsOptions contains the optional parameters for the Client.getChatCompletions method.
func (client *Client) getChatCompletions(ctx context.Context, body ChatCompletionsOptions, options *GetChatCompletionsOptions) (GetChatCompletionsResponse, error) {
	var err error
	req, err := client.getChatCompletionsCreateRequest(ctx, body, options)
	if err != nil {
		return GetChatCompletionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetChatCompletionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = client.newError(httpResp)
		return GetChatCompletionsResponse{}, err
	}
	resp, err := client.getChatCompletionsHandleResponse(httpResp)
	return resp, err
}

// getChatCompletionsCreateRequest creates the getChatCompletions request.
func (client *Client) getChatCompletionsCreateRequest(ctx context.Context, body ChatCompletionsOptions, options *GetChatCompletionsOptions) (*policy.Request, error) {
	urlPath := "chat/completions"
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.formatURL(urlPath, getDeployment(body)))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getChatCompletionsHandleResponse handles the getChatCompletions response.
func (client *Client) getChatCompletionsHandleResponse(resp *http.Response) (GetChatCompletionsResponse, error) {
	result := GetChatCompletionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChatCompletions); err != nil {
		return GetChatCompletionsResponse{}, err
	}
	return result, nil
}

// getChatCompletionsWithAzureExtensions - Gets chat completions for the provided chat messages. This is an Azure-specific
// version of chat completions that supports integration with configured data sources and other augmentations to the base
// chat completions capabilities.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - options - GetChatCompletionsWithAzureExtensionsOptions contains the optional parameters for the Client.GetChatCompletionsWithAzureExtensions
//     method.
func (client *Client) getChatCompletionsWithAzureExtensions(ctx context.Context, body ChatCompletionsOptions, options *GetChatCompletionsWithAzureExtensionsOptions) (GetChatCompletionsWithAzureExtensionsResponse, error) {
	var err error
	req, err := client.getChatCompletionsWithAzureExtensionsCreateRequest(ctx, body, options)
	if err != nil {
		return GetChatCompletionsWithAzureExtensionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetChatCompletionsWithAzureExtensionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = client.newError(httpResp)
		return GetChatCompletionsWithAzureExtensionsResponse{}, err
	}
	resp, err := client.getChatCompletionsWithAzureExtensionsHandleResponse(httpResp)
	return resp, err
}

// getChatCompletionsWithAzureExtensionsCreateRequest creates the getChatCompletionsWithAzureExtensions request.
func (client *Client) getChatCompletionsWithAzureExtensionsCreateRequest(ctx context.Context, body ChatCompletionsOptions, options *GetChatCompletionsWithAzureExtensionsOptions) (*policy.Request, error) {
	urlPath := "extensions/chat/completions"
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.formatURL(urlPath, getDeployment(body)))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getChatCompletionsWithAzureExtensionsHandleResponse handles the getChatCompletionsWithAzureExtensions response.
func (client *Client) getChatCompletionsWithAzureExtensionsHandleResponse(resp *http.Response) (GetChatCompletionsWithAzureExtensionsResponse, error) {
	result := GetChatCompletionsWithAzureExtensionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ChatCompletions); err != nil {
		return GetChatCompletionsWithAzureExtensionsResponse{}, err
	}
	return result, nil
}

// GetCompletions - Gets completions for the provided input prompts. Completions support a wide variety of tasks and generate
// text that continues from or "completes" provided prompt data.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - options - GetCompletionsOptions contains the optional parameters for the Client.GetCompletions method.
func (client *Client) GetCompletions(ctx context.Context, body CompletionsOptions, options *GetCompletionsOptions) (GetCompletionsResponse, error) {
	var err error
	req, err := client.getCompletionsCreateRequest(ctx, body, options)
	if err != nil {
		return GetCompletionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetCompletionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = client.newError(httpResp)
		return GetCompletionsResponse{}, err
	}
	resp, err := client.getCompletionsHandleResponse(httpResp)
	return resp, err
}

// getCompletionsCreateRequest creates the GetCompletions request.
func (client *Client) getCompletionsCreateRequest(ctx context.Context, body CompletionsOptions, options *GetCompletionsOptions) (*policy.Request, error) {
	urlPath := "completions"
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.formatURL(urlPath, getDeployment(body)))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getCompletionsHandleResponse handles the GetCompletions response.
func (client *Client) getCompletionsHandleResponse(resp *http.Response) (GetCompletionsResponse, error) {
	result := GetCompletionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Completions); err != nil {
		return GetCompletionsResponse{}, err
	}
	return result, nil
}

// GetEmbeddings - Return the embeddings for a given prompt.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - options - GetEmbeddingsOptions contains the optional parameters for the Client.GetEmbeddings method.
func (client *Client) GetEmbeddings(ctx context.Context, body EmbeddingsOptions, options *GetEmbeddingsOptions) (GetEmbeddingsResponse, error) {
	var err error
	req, err := client.getEmbeddingsCreateRequest(ctx, body, options)
	if err != nil {
		return GetEmbeddingsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetEmbeddingsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = client.newError(httpResp)
		return GetEmbeddingsResponse{}, err
	}
	resp, err := client.getEmbeddingsHandleResponse(httpResp)
	return resp, err
}

// getEmbeddingsCreateRequest creates the GetEmbeddings request.
func (client *Client) getEmbeddingsCreateRequest(ctx context.Context, body EmbeddingsOptions, options *GetEmbeddingsOptions) (*policy.Request, error) {
	urlPath := "embeddings"
	req, err := runtime.NewRequest(ctx, http.MethodPost, client.formatURL(urlPath, getDeployment(body)))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// getEmbeddingsHandleResponse handles the GetEmbeddings response.
func (client *Client) getEmbeddingsHandleResponse(resp *http.Response) (GetEmbeddingsResponse, error) {
	result := GetEmbeddingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Embeddings); err != nil {
		return GetEmbeddingsResponse{}, err
	}
	return result, nil
}
