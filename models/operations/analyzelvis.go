// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/const-cloudinary/analysis-go/models/components"
)

type AnalyzeLvisGlobals struct {
	// The name of your Cloudinary cloud
	CloudName *string `pathParam:"style=simple,explode=false,name=cloud_name"`
}

func (o *AnalyzeLvisGlobals) GetCloudName() *string {
	if o == nil {
		return nil
	}
	return o.CloudName
}

type AnalyzeLvisRequest struct {
	// The name of your Cloudinary cloud
	CloudName *string `pathParam:"style=simple,explode=false,name=cloud_name"`
	// A JSON object containing request parameters
	BaseAnalyzeRequest components.BaseAnalyzeRequest `request:"mediaType=application/json"`
}

func (o *AnalyzeLvisRequest) GetCloudName() *string {
	if o == nil {
		return nil
	}
	return o.CloudName
}

func (o *AnalyzeLvisRequest) GetBaseAnalyzeRequest() components.BaseAnalyzeRequest {
	if o == nil {
		return components.BaseAnalyzeRequest{}
	}
	return o.BaseAnalyzeRequest
}

type AnalyzeLvisResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Analysis succeeded
	AnalyzeLvisResponse *components.AnalyzeLvisResponse
}

func (o *AnalyzeLvisResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AnalyzeLvisResponse) GetAnalyzeLvisResponse() *components.AnalyzeLvisResponse {
	if o == nil {
		return nil
	}
	return o.AnalyzeLvisResponse
}
