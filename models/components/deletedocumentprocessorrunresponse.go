// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DeleteDocumentProcessorRunResponse struct {
	// Unique identifier of the document processor run.
	ID string `json:"id"`
}

func (o *DeleteDocumentProcessorRunResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
