// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AnalyzeCocoResponseData struct {
	Entity   *string           `json:"entity,omitempty"`
	Analysis *CocoAnalysisData `json:"analysis,omitempty"`
}

func (o *AnalyzeCocoResponseData) GetEntity() *string {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *AnalyzeCocoResponseData) GetAnalysis() *CocoAnalysisData {
	if o == nil {
		return nil
	}
	return o.Analysis
}

type AnalyzeCocoResponse struct {
	Data      *AnalyzeCocoResponseData `json:"data,omitempty"`
	Limits    *Limits                  `json:"limits,omitempty"`
	RequestID *string                  `json:"request_id,omitempty"`
}

func (o *AnalyzeCocoResponse) GetData() *AnalyzeCocoResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AnalyzeCocoResponse) GetLimits() *Limits {
	if o == nil {
		return nil
	}
	return o.Limits
}

func (o *AnalyzeCocoResponse) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}
