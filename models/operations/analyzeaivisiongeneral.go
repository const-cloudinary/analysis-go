// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/const-cloudinary/analysis-go/models/components"
)

type AnalyzeAiVisionGeneralGlobals struct {
	// The name of your Cloudinary cloud
	CloudName *string `pathParam:"style=simple,explode=false,name=cloud_name"`
}

func (o *AnalyzeAiVisionGeneralGlobals) GetCloudName() *string {
	if o == nil {
		return nil
	}
	return o.CloudName
}

type AnalyzeAiVisionGeneralRequest struct {
	// The name of your Cloudinary cloud
	CloudName *string `pathParam:"style=simple,explode=false,name=cloud_name"`
	// A JSON object containing request parameters
	AnalyzeAIVisionGeneralRequest components.AnalyzeAIVisionGeneralRequest `request:"mediaType=application/json"`
}

func (o *AnalyzeAiVisionGeneralRequest) GetCloudName() *string {
	if o == nil {
		return nil
	}
	return o.CloudName
}

func (o *AnalyzeAiVisionGeneralRequest) GetAnalyzeAIVisionGeneralRequest() components.AnalyzeAIVisionGeneralRequest {
	if o == nil {
		return components.AnalyzeAIVisionGeneralRequest{}
	}
	return o.AnalyzeAIVisionGeneralRequest
}

type AnalyzeAiVisionGeneralResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Analysis succeeded
	AnalyzeAIVisionGeneralResponse *components.AnalyzeAIVisionGeneralResponse
}

func (o *AnalyzeAiVisionGeneralResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AnalyzeAiVisionGeneralResponse) GetAnalyzeAIVisionGeneralResponse() *components.AnalyzeAIVisionGeneralResponse {
	if o == nil {
		return nil
	}
	return o.AnalyzeAIVisionGeneralResponse
}
