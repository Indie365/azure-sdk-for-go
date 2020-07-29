package training

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.2/customvision/training"

// BoundingBox ...
type BoundingBox struct {
	Left   *float64 `json:"left,omitempty"`
	Top    *float64 `json:"top,omitempty"`
	Width  *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
}

// Domain ...
type Domain struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - READ-ONLY
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Possible values include: 'Classification', 'ObjectDetection'
	Type DomainType `json:"type,omitempty"`
	// Exportable - READ-ONLY
	Exportable *bool `json:"exportable,omitempty"`
	// Enabled - READ-ONLY
	Enabled *bool `json:"enabled,omitempty"`
}

// MarshalJSON is the custom marshaler for Domain.
func (d Domain) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Export ...
type Export struct {
	autorest.Response `json:"-"`
	// Platform - READ-ONLY; Platform of the export. Possible values include: 'CoreML', 'TensorFlow', 'DockerFile', 'ONNX'
	Platform ExportPlatformModel `json:"platform,omitempty"`
	// Status - READ-ONLY; Status of the export. Possible values include: 'Exporting', 'Failed', 'Done'
	Status ExportStatusModel `json:"status,omitempty"`
	// DownloadURI - READ-ONLY; URI used to download the model.
	DownloadURI *string `json:"downloadUri,omitempty"`
	// Flavor - READ-ONLY; Flavor of the export. Possible values include: 'Linux', 'Windows', 'ONNX10', 'ONNX12'
	Flavor ExportFlavorModel `json:"flavor,omitempty"`
	// NewerVersionAvailable - READ-ONLY; Indicates an updated version of the export package is available and should be re-exported for the latest changes.
	NewerVersionAvailable *bool `json:"newerVersionAvailable,omitempty"`
}

// MarshalJSON is the custom marshaler for Export.
func (e Export) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Image image model to be sent as JSON.
type Image struct {
	// ID - READ-ONLY; Id of the image.
	ID *uuid.UUID `json:"id,omitempty"`
	// Created - READ-ONLY; Date the image was created.
	Created *date.Time `json:"created,omitempty"`
	// Width - READ-ONLY; Width of the image.
	Width *int32 `json:"width,omitempty"`
	// Height - READ-ONLY; Height of the image.
	Height *int32 `json:"height,omitempty"`
	// ResizedImageURI - READ-ONLY; The URI to the (resized) image used for training.
	ResizedImageURI *string `json:"resizedImageUri,omitempty"`
	// ThumbnailURI - READ-ONLY; The URI to the thumbnail of the original image.
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
	// OriginalImageURI - READ-ONLY; The URI to the original uploaded image.
	OriginalImageURI *string `json:"originalImageUri,omitempty"`
	// Tags - READ-ONLY; Tags associated with this image.
	Tags *[]ImageTag `json:"tags,omitempty"`
	// Regions - READ-ONLY; Regions associated with this image.
	Regions *[]ImageRegion `json:"regions,omitempty"`
}

// MarshalJSON is the custom marshaler for Image.
func (i Image) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageCreateResult ...
type ImageCreateResult struct {
	// SourceURL - READ-ONLY; Source URL of the image.
	SourceURL *string `json:"sourceUrl,omitempty"`
	// Status - READ-ONLY; Status of the image creation. Possible values include: 'OK', 'OKDuplicate', 'ErrorSource', 'ErrorImageFormat', 'ErrorImageSize', 'ErrorStorage', 'ErrorLimitExceed', 'ErrorTagLimitExceed', 'ErrorRegionLimitExceed', 'ErrorUnknown', 'ErrorNegativeAndRegularTagOnSameImage'
	Status ImageCreateStatus `json:"status,omitempty"`
	// Image - READ-ONLY; The image.
	Image *Image `json:"image,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageCreateResult.
func (icr ImageCreateResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageCreateSummary ...
type ImageCreateSummary struct {
	autorest.Response `json:"-"`
	// IsBatchSuccessful - READ-ONLY; True if all of the images in the batch were created successfully, otherwise false.
	IsBatchSuccessful *bool `json:"isBatchSuccessful,omitempty"`
	// Images - READ-ONLY; List of the image creation results.
	Images *[]ImageCreateResult `json:"images,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageCreateSummary.
func (ics ImageCreateSummary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageFileCreateBatch ...
type ImageFileCreateBatch struct {
	Images *[]ImageFileCreateEntry `json:"images,omitempty"`
	TagIds *[]uuid.UUID            `json:"tagIds,omitempty"`
}

// ImageFileCreateEntry ...
type ImageFileCreateEntry struct {
	Name     *string      `json:"name,omitempty"`
	Contents *[]byte      `json:"contents,omitempty"`
	TagIds   *[]uuid.UUID `json:"tagIds,omitempty"`
	Regions  *[]Region    `json:"regions,omitempty"`
}

// ImageIDCreateBatch ...
type ImageIDCreateBatch struct {
	Images *[]ImageIDCreateEntry `json:"images,omitempty"`
	TagIds *[]uuid.UUID          `json:"tagIds,omitempty"`
}

// ImageIDCreateEntry ...
type ImageIDCreateEntry struct {
	ID      *uuid.UUID   `json:"id,omitempty"`
	TagIds  *[]uuid.UUID `json:"tagIds,omitempty"`
	Regions *[]Region    `json:"regions,omitempty"`
}

// ImagePerformance image performance model.
type ImagePerformance struct {
	// Predictions - READ-ONLY
	Predictions *[]Prediction `json:"predictions,omitempty"`
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// Width - READ-ONLY
	Width *int32 `json:"width,omitempty"`
	// Height - READ-ONLY
	Height *int32 `json:"height,omitempty"`
	// ImageURI - READ-ONLY
	ImageURI *string `json:"imageUri,omitempty"`
	// ThumbnailURI - READ-ONLY
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
	// Tags - READ-ONLY
	Tags *[]ImageTag `json:"tags,omitempty"`
	// Regions - READ-ONLY
	Regions *[]ImageRegion `json:"regions,omitempty"`
}

// MarshalJSON is the custom marshaler for ImagePerformance.
func (IP ImagePerformance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImagePrediction ...
type ImagePrediction struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Project - READ-ONLY
	Project *uuid.UUID `json:"project,omitempty"`
	// Iteration - READ-ONLY
	Iteration *uuid.UUID `json:"iteration,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// Predictions - READ-ONLY
	Predictions *[]Prediction `json:"predictions,omitempty"`
}

// MarshalJSON is the custom marshaler for ImagePrediction.
func (IP ImagePrediction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageRegion ...
type ImageRegion struct {
	// RegionID - READ-ONLY
	RegionID *uuid.UUID `json:"regionId,omitempty"`
	// TagName - READ-ONLY
	TagName *string `json:"tagName,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// TagID - Id of the tag associated with this region.
	TagID  *uuid.UUID `json:"tagId,omitempty"`
	Left   *float64   `json:"left,omitempty"`
	Top    *float64   `json:"top,omitempty"`
	Width  *float64   `json:"width,omitempty"`
	Height *float64   `json:"height,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageRegion.
func (ir ImageRegion) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ir.TagID != nil {
		objectMap["tagId"] = ir.TagID
	}
	if ir.Left != nil {
		objectMap["left"] = ir.Left
	}
	if ir.Top != nil {
		objectMap["top"] = ir.Top
	}
	if ir.Width != nil {
		objectMap["width"] = ir.Width
	}
	if ir.Height != nil {
		objectMap["height"] = ir.Height
	}
	return json.Marshal(objectMap)
}

// ImageRegionCreateBatch batch of image region information to create.
type ImageRegionCreateBatch struct {
	Regions *[]ImageRegionCreateEntry `json:"regions,omitempty"`
}

// ImageRegionCreateEntry entry associating a region to an image.
type ImageRegionCreateEntry struct {
	// ImageID - Id of the image.
	ImageID *uuid.UUID `json:"imageId,omitempty"`
	// TagID - Id of the tag associated with this region.
	TagID  *uuid.UUID `json:"tagId,omitempty"`
	Left   *float64   `json:"left,omitempty"`
	Top    *float64   `json:"top,omitempty"`
	Width  *float64   `json:"width,omitempty"`
	Height *float64   `json:"height,omitempty"`
}

// ImageRegionCreateResult ...
type ImageRegionCreateResult struct {
	// ImageID - READ-ONLY
	ImageID *uuid.UUID `json:"imageId,omitempty"`
	// RegionID - READ-ONLY
	RegionID *uuid.UUID `json:"regionId,omitempty"`
	// TagName - READ-ONLY
	TagName *string `json:"tagName,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// TagID - Id of the tag associated with this region.
	TagID  *uuid.UUID `json:"tagId,omitempty"`
	Left   *float64   `json:"left,omitempty"`
	Top    *float64   `json:"top,omitempty"`
	Width  *float64   `json:"width,omitempty"`
	Height *float64   `json:"height,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageRegionCreateResult.
func (ircr ImageRegionCreateResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ircr.TagID != nil {
		objectMap["tagId"] = ircr.TagID
	}
	if ircr.Left != nil {
		objectMap["left"] = ircr.Left
	}
	if ircr.Top != nil {
		objectMap["top"] = ircr.Top
	}
	if ircr.Width != nil {
		objectMap["width"] = ircr.Width
	}
	if ircr.Height != nil {
		objectMap["height"] = ircr.Height
	}
	return json.Marshal(objectMap)
}

// ImageRegionCreateSummary ...
type ImageRegionCreateSummary struct {
	autorest.Response `json:"-"`
	Created           *[]ImageRegionCreateResult `json:"created,omitempty"`
	Duplicated        *[]ImageRegionCreateEntry  `json:"duplicated,omitempty"`
	Exceeded          *[]ImageRegionCreateEntry  `json:"exceeded,omitempty"`
}

// ImageRegionProposal ...
type ImageRegionProposal struct {
	autorest.Response `json:"-"`
	// ProjectID - READ-ONLY
	ProjectID *uuid.UUID `json:"projectId,omitempty"`
	// ImageID - READ-ONLY
	ImageID *uuid.UUID `json:"imageId,omitempty"`
	// Proposals - READ-ONLY
	Proposals *[]RegionProposal `json:"proposals,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageRegionProposal.
func (irp ImageRegionProposal) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageTag ...
type ImageTag struct {
	// TagID - READ-ONLY
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// TagName - READ-ONLY
	TagName *string `json:"tagName,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
}

// MarshalJSON is the custom marshaler for ImageTag.
func (it ImageTag) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ImageTagCreateBatch batch of image tags.
type ImageTagCreateBatch struct {
	// Tags - Image Tag entries to include in this batch.
	Tags *[]ImageTagCreateEntry `json:"tags,omitempty"`
}

// ImageTagCreateEntry entry associating a tag to an image.
type ImageTagCreateEntry struct {
	// ImageID - Id of the image.
	ImageID *uuid.UUID `json:"imageId,omitempty"`
	// TagID - Id of the tag.
	TagID *uuid.UUID `json:"tagId,omitempty"`
}

// ImageTagCreateSummary ...
type ImageTagCreateSummary struct {
	autorest.Response `json:"-"`
	Created           *[]ImageTagCreateEntry `json:"created,omitempty"`
	Duplicated        *[]ImageTagCreateEntry `json:"duplicated,omitempty"`
	Exceeded          *[]ImageTagCreateEntry `json:"exceeded,omitempty"`
}

// ImageURL ...
type ImageURL struct {
	URL *string `json:"url,omitempty"`
}

// ImageURLCreateBatch ...
type ImageURLCreateBatch struct {
	Images *[]ImageURLCreateEntry `json:"images,omitempty"`
	TagIds *[]uuid.UUID           `json:"tagIds,omitempty"`
}

// ImageURLCreateEntry ...
type ImageURLCreateEntry struct {
	URL     *string      `json:"url,omitempty"`
	TagIds  *[]uuid.UUID `json:"tagIds,omitempty"`
	Regions *[]Region    `json:"regions,omitempty"`
}

// Int32 ...
type Int32 struct {
	autorest.Response `json:"-"`
	Value             *int32 `json:"value,omitempty"`
}

// Iteration iteration model to be sent over JSON.
type Iteration struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Gets the id of the iteration.
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - Gets or sets the name of the iteration.
	Name *string `json:"name,omitempty"`
	// IsDefault - Gets or sets a value indicating whether the iteration is the default iteration for the project.
	IsDefault *bool `json:"isDefault,omitempty"`
	// Status - READ-ONLY; Gets the current iteration status.
	Status *string `json:"status,omitempty"`
	// Created - READ-ONLY; Gets the time this iteration was completed.
	Created *date.Time `json:"created,omitempty"`
	// LastModified - READ-ONLY; Gets the time this iteration was last modified.
	LastModified *date.Time `json:"lastModified,omitempty"`
	// TrainedAt - READ-ONLY; Gets the time this iteration was last modified.
	TrainedAt *date.Time `json:"trainedAt,omitempty"`
	// ProjectID - READ-ONLY; Gets The project id. of the iteration.
	ProjectID *uuid.UUID `json:"projectId,omitempty"`
	// Exportable - READ-ONLY; Whether the iteration can be exported to another format for download.
	Exportable *bool `json:"exportable,omitempty"`
	// DomainID - READ-ONLY; Get or sets a guid of the domain the iteration has been trained on.
	DomainID *uuid.UUID `json:"domainId,omitempty"`
	// ClassificationType - READ-ONLY; Gets the classification type of the project. Possible values include: 'Multiclass', 'Multilabel'
	ClassificationType Classifier `json:"classificationType,omitempty"`
}

// MarshalJSON is the custom marshaler for Iteration.
func (i Iteration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if i.Name != nil {
		objectMap["name"] = i.Name
	}
	if i.IsDefault != nil {
		objectMap["isDefault"] = i.IsDefault
	}
	return json.Marshal(objectMap)
}

// IterationPerformance represents the detailed performance data for a trained iteration.
type IterationPerformance struct {
	autorest.Response `json:"-"`
	// PerTagPerformance - READ-ONLY; Gets the per-tag performance details for this iteration.
	PerTagPerformance *[]TagPerformance `json:"perTagPerformance,omitempty"`
	// Precision - READ-ONLY; Gets the precision.
	Precision *float64 `json:"precision,omitempty"`
	// PrecisionStdDeviation - READ-ONLY; Gets the standard deviation for the precision.
	PrecisionStdDeviation *float64 `json:"precisionStdDeviation,omitempty"`
	// Recall - READ-ONLY; Gets the recall.
	Recall *float64 `json:"recall,omitempty"`
	// RecallStdDeviation - READ-ONLY; Gets the standard deviation for the recall.
	RecallStdDeviation *float64 `json:"recallStdDeviation,omitempty"`
	// AveragePrecision - READ-ONLY; Gets the average precision when applicable.
	AveragePrecision *float64 `json:"averagePrecision,omitempty"`
}

// MarshalJSON is the custom marshaler for IterationPerformance.
func (IP IterationPerformance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ListDomain ...
type ListDomain struct {
	autorest.Response `json:"-"`
	Value             *[]Domain `json:"value,omitempty"`
}

// ListExport ...
type ListExport struct {
	autorest.Response `json:"-"`
	Value             *[]Export `json:"value,omitempty"`
}

// ListImage ...
type ListImage struct {
	autorest.Response `json:"-"`
	Value             *[]Image `json:"value,omitempty"`
}

// ListImagePerformance ...
type ListImagePerformance struct {
	autorest.Response `json:"-"`
	Value             *[]ImagePerformance `json:"value,omitempty"`
}

// ListIteration ...
type ListIteration struct {
	autorest.Response `json:"-"`
	Value             *[]Iteration `json:"value,omitempty"`
}

// ListProject ...
type ListProject struct {
	autorest.Response `json:"-"`
	Value             *[]Project `json:"value,omitempty"`
}

// ListTag ...
type ListTag struct {
	autorest.Response `json:"-"`
	Value             *[]Tag `json:"value,omitempty"`
}

// Prediction ...
type Prediction struct {
	// Probability - READ-ONLY
	Probability *float64 `json:"probability,omitempty"`
	// TagID - READ-ONLY
	TagID *uuid.UUID `json:"tagId,omitempty"`
	// TagName - READ-ONLY
	TagName *string `json:"tagName,omitempty"`
	// BoundingBox - READ-ONLY
	BoundingBox *BoundingBox `json:"boundingBox,omitempty"`
}

// MarshalJSON is the custom marshaler for Prediction.
func (p Prediction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// PredictionQueryResult ...
type PredictionQueryResult struct {
	autorest.Response `json:"-"`
	// Token - READ-ONLY
	Token *PredictionQueryToken `json:"token,omitempty"`
	// Results - READ-ONLY
	Results *[]StoredImagePrediction `json:"results,omitempty"`
}

// MarshalJSON is the custom marshaler for PredictionQueryResult.
func (pqr PredictionQueryResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// PredictionQueryTag ...
type PredictionQueryTag struct {
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// MinThreshold - READ-ONLY
	MinThreshold *float64 `json:"minThreshold,omitempty"`
	// MaxThreshold - READ-ONLY
	MaxThreshold *float64 `json:"maxThreshold,omitempty"`
}

// MarshalJSON is the custom marshaler for PredictionQueryTag.
func (pqt PredictionQueryTag) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// PredictionQueryToken ...
type PredictionQueryToken struct {
	Session      *string `json:"session,omitempty"`
	Continuation *string `json:"continuation,omitempty"`
	MaxCount     *int32  `json:"maxCount,omitempty"`
	// OrderBy - Possible values include: 'Newest', 'Oldest', 'Suggested'
	OrderBy     OrderBy               `json:"orderBy,omitempty"`
	Tags        *[]PredictionQueryTag `json:"tags,omitempty"`
	IterationID *uuid.UUID            `json:"iterationId,omitempty"`
	StartTime   *date.Time            `json:"startTime,omitempty"`
	EndTime     *date.Time            `json:"endTime,omitempty"`
	Application *string               `json:"application,omitempty"`
}

// Project represents a project.
type Project struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Gets The project id.
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - Gets or sets the name of the project.
	Name *string `json:"name,omitempty"`
	// Description - Gets or sets the description of the project.
	Description *string `json:"description,omitempty"`
	// Settings - Gets or sets the project settings.
	Settings *ProjectSettings `json:"settings,omitempty"`
	// Created - READ-ONLY; Gets the date this project was created.
	Created *date.Time `json:"created,omitempty"`
	// LastModified - READ-ONLY; Gets the date this project was last modified.
	LastModified *date.Time `json:"lastModified,omitempty"`
	// ThumbnailURI - READ-ONLY; Gets the thumbnail url representing the project.
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
}

// MarshalJSON is the custom marshaler for Project.
func (p Project) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.Name != nil {
		objectMap["name"] = p.Name
	}
	if p.Description != nil {
		objectMap["description"] = p.Description
	}
	if p.Settings != nil {
		objectMap["settings"] = p.Settings
	}
	return json.Marshal(objectMap)
}

// ProjectSettings represents settings associated with a project.
type ProjectSettings struct {
	// DomainID - Gets or sets the id of the Domain to use with this project.
	DomainID *uuid.UUID `json:"domainId,omitempty"`
	// ClassificationType - Gets or sets the classification type of the project. Possible values include: 'Multiclass', 'Multilabel'
	ClassificationType Classifier `json:"classificationType,omitempty"`
}

// Region ...
type Region struct {
	// TagID - Id of the tag associated with this region.
	TagID  *uuid.UUID `json:"tagId,omitempty"`
	Left   *float64   `json:"left,omitempty"`
	Top    *float64   `json:"top,omitempty"`
	Width  *float64   `json:"width,omitempty"`
	Height *float64   `json:"height,omitempty"`
}

// RegionProposal ...
type RegionProposal struct {
	// Confidence - READ-ONLY
	Confidence *float64 `json:"confidence,omitempty"`
	// BoundingBox - READ-ONLY
	BoundingBox *BoundingBox `json:"boundingBox,omitempty"`
}

// MarshalJSON is the custom marshaler for RegionProposal.
func (rp RegionProposal) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// StoredImagePrediction result of an image classification request.
type StoredImagePrediction struct {
	// ResizedImageURI - READ-ONLY; The URI to the (resized) prediction image.
	ResizedImageURI *string `json:"resizedImageUri,omitempty"`
	// ThumbnailURI - READ-ONLY; The URI to the thumbnail of the original prediction image.
	ThumbnailURI *string `json:"thumbnailUri,omitempty"`
	// OriginalImageURI - READ-ONLY; The URI to the original prediction image.
	OriginalImageURI *string `json:"originalImageUri,omitempty"`
	// Domain - READ-ONLY; Domain used for the prediction.
	Domain *uuid.UUID `json:"domain,omitempty"`
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Project - READ-ONLY
	Project *uuid.UUID `json:"project,omitempty"`
	// Iteration - READ-ONLY
	Iteration *uuid.UUID `json:"iteration,omitempty"`
	// Created - READ-ONLY
	Created *date.Time `json:"created,omitempty"`
	// Predictions - READ-ONLY
	Predictions *[]Prediction `json:"predictions,omitempty"`
}

// MarshalJSON is the custom marshaler for StoredImagePrediction.
func (sip StoredImagePrediction) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Tag represents a Tag.
type Tag struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Gets the Tag ID.
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - Gets or sets the name of the tag.
	Name *string `json:"name,omitempty"`
	// Description - Gets or sets the description of the tag.
	Description *string `json:"description,omitempty"`
	// Type - Gets or sets the type of the tag. Possible values include: 'Regular', 'Negative'
	Type TagType `json:"type,omitempty"`
	// ImageCount - READ-ONLY; Gets the number of images with this tag.
	ImageCount *int32 `json:"imageCount,omitempty"`
}

// MarshalJSON is the custom marshaler for Tag.
func (t Tag) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if t.Name != nil {
		objectMap["name"] = t.Name
	}
	if t.Description != nil {
		objectMap["description"] = t.Description
	}
	if t.Type != "" {
		objectMap["type"] = t.Type
	}
	return json.Marshal(objectMap)
}

// TagPerformance represents performance data for a particular tag in a trained iteration.
type TagPerformance struct {
	// ID - READ-ONLY
	ID *uuid.UUID `json:"id,omitempty"`
	// Name - READ-ONLY
	Name *string `json:"name,omitempty"`
	// Precision - READ-ONLY; Gets the precision.
	Precision *float64 `json:"precision,omitempty"`
	// PrecisionStdDeviation - READ-ONLY; Gets the standard deviation for the precision.
	PrecisionStdDeviation *float64 `json:"precisionStdDeviation,omitempty"`
	// Recall - READ-ONLY; Gets the recall.
	Recall *float64 `json:"recall,omitempty"`
	// RecallStdDeviation - READ-ONLY; Gets the standard deviation for the recall.
	RecallStdDeviation *float64 `json:"recallStdDeviation,omitempty"`
	// AveragePrecision - READ-ONLY; Gets the average precision when applicable.
	AveragePrecision *float64 `json:"averagePrecision,omitempty"`
}

// MarshalJSON is the custom marshaler for TagPerformance.
func (tp TagPerformance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}
