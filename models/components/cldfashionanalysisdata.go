// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CldFashionAnalysisDataTags struct {
}

type CldFashionAnalysisData struct {
	Tags         CldFashionAnalysisDataTags `json:"tags"`
	ModelVersion *int64                     `json:"model_version,omitempty"`
}

func (o *CldFashionAnalysisData) GetTags() CldFashionAnalysisDataTags {
	if o == nil {
		return CldFashionAnalysisDataTags{}
	}
	return o.Tags
}

func (o *CldFashionAnalysisData) GetModelVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.ModelVersion
}
