// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CldTextAnalysisDataTags struct {
}

type CldTextAnalysisData struct {
	Tags         CldTextAnalysisDataTags `json:"tags"`
	ModelVersion *int64                  `json:"model_version,omitempty"`
}

func (o *CldTextAnalysisData) GetTags() CldTextAnalysisDataTags {
	if o == nil {
		return CldTextAnalysisDataTags{}
	}
	return o.Tags
}

func (o *CldTextAnalysisData) GetModelVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.ModelVersion
}
