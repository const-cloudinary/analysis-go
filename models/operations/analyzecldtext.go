// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/const-cloudinary/analysis-go/models/components"
)

type AnalyzeCldTextGlobals struct {
	// The name of your Cloudinary cloud
	CloudName *string `pathParam:"style=simple,explode=false,name=cloud_name"`
}

func (o *AnalyzeCldTextGlobals) GetCloudName() *string {
	if o == nil {
		return nil
	}
	return o.CloudName
}

type AnalyzeCldTextRequest struct {
	// The name of your Cloudinary cloud
	CloudName *string `pathParam:"style=simple,explode=false,name=cloud_name"`
	// A JSON object containing request parameters
	BaseAnalyzeRequest components.BaseAnalyzeRequest `request:"mediaType=application/json"`
}

func (o *AnalyzeCldTextRequest) GetCloudName() *string {
	if o == nil {
		return nil
	}
	return o.CloudName
}

func (o *AnalyzeCldTextRequest) GetBaseAnalyzeRequest() components.BaseAnalyzeRequest {
	if o == nil {
		return components.BaseAnalyzeRequest{}
	}
	return o.BaseAnalyzeRequest
}

type AnalyzeCldTextResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Analysis succeeded
	AnalyzeCldTextResponse *components.AnalyzeCldTextResponse
}

func (o *AnalyzeCldTextResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AnalyzeCldTextResponse) GetAnalyzeCldTextResponse() *components.AnalyzeCldTextResponse {
	if o == nil {
		return nil
	}
	return o.AnalyzeCldTextResponse
}
