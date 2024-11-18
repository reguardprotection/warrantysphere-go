// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateDocumentProcessorRunBody struct {
	// Unique identifier of the document processor.
	ProcessorID string `json:"processorId"`
	// Unique identifier of the document.
	DocumentID string `json:"documentId"`
}

func (o *CreateDocumentProcessorRunBody) GetProcessorID() string {
	if o == nil {
		return ""
	}
	return o.ProcessorID
}

func (o *CreateDocumentProcessorRunBody) GetDocumentID() string {
	if o == nil {
		return ""
	}
	return o.DocumentID
}
