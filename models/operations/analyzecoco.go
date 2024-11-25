// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/const-cloudinary/analysis-go/models/components"
)

type AnalyzeCocoRequest struct {
	// The name of your Cloudinary cloud
	CloudName string `pathParam:"style=simple,explode=false,name=cloud_name"`
	// A JSON object containing request parameters
	BaseAnalyzeRequest components.BaseAnalyzeRequest `request:"mediaType=application/json"`
}

func (o *AnalyzeCocoRequest) GetCloudName() string {
	if o == nil {
		return ""
	}
	return o.CloudName
}

func (o *AnalyzeCocoRequest) GetBaseAnalyzeRequest() components.BaseAnalyzeRequest {
	if o == nil {
		return components.BaseAnalyzeRequest{}
	}
	return o.BaseAnalyzeRequest
}

type AnalyzeCocoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Analysis succeeded
	AnalyzeCocoResponse *components.AnalyzeCocoResponse
}

func (o *AnalyzeCocoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AnalyzeCocoResponse) GetAnalyzeCocoResponse() *components.AnalyzeCocoResponse {
	if o == nil {
		return nil
	}
	return o.AnalyzeCocoResponse
}
