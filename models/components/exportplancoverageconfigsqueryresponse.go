// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ExportPlanCoverageConfigsQueryResponse struct {
	// The exported plan coverage config data in csv format
	Data string `json:"data"`
}

func (o *ExportPlanCoverageConfigsQueryResponse) GetData() string {
	if o == nil {
		return ""
	}
	return o.Data
}
