// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/cloudinary/analysis-go/models/components"
)

type AnalyzeAiVisionTaggingRequest struct {
	// The name of your Cloudinary cloud
	CloudName string `pathParam:"style=simple,explode=false,name=cloud_name"`
	// A JSON object containing request parameters
	AnalyzeAIVisionTaggingRequest components.AnalyzeAIVisionTaggingRequest `request:"mediaType=application/json"`
}

func (o *AnalyzeAiVisionTaggingRequest) GetCloudName() string {
	if o == nil {
		return ""
	}
	return o.CloudName
}

func (o *AnalyzeAiVisionTaggingRequest) GetAnalyzeAIVisionTaggingRequest() components.AnalyzeAIVisionTaggingRequest {
	if o == nil {
		return components.AnalyzeAIVisionTaggingRequest{}
	}
	return o.AnalyzeAIVisionTaggingRequest
}

type AnalyzeAiVisionTaggingResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Analysis succeeded
	AnalyzeAIVisionTaggingResponse *components.AnalyzeAIVisionTaggingResponse
}

func (o *AnalyzeAiVisionTaggingResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AnalyzeAiVisionTaggingResponse) GetAnalyzeAIVisionTaggingResponse() *components.AnalyzeAIVisionTaggingResponse {
	if o == nil {
		return nil
	}
	return o.AnalyzeAIVisionTaggingResponse
}
