// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AnalyzeGoogleTaggingResponseData struct {
	Entity   *string                    `json:"entity,omitempty"`
	Analysis *GoogleTaggingAnalysisData `json:"analysis,omitempty"`
}

func (o *AnalyzeGoogleTaggingResponseData) GetEntity() *string {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *AnalyzeGoogleTaggingResponseData) GetAnalysis() *GoogleTaggingAnalysisData {
	if o == nil {
		return nil
	}
	return o.Analysis
}

type AnalyzeGoogleTaggingResponse struct {
	Data      *AnalyzeGoogleTaggingResponseData `json:"data,omitempty"`
	Limits    *Limits                           `json:"limits,omitempty"`
	RequestID *string                           `json:"request_id,omitempty"`
}

func (o *AnalyzeGoogleTaggingResponse) GetData() *AnalyzeGoogleTaggingResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AnalyzeGoogleTaggingResponse) GetLimits() *Limits {
	if o == nil {
		return nil
	}
	return o.Limits
}

func (o *AnalyzeGoogleTaggingResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}
