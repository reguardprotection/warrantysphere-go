// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type DocumentProcessorRunsDeleteRequest struct {
	// Unique identifier of the document processor run.
	RunID string `pathParam:"style=simple,explode=false,name=run_id"`
}

func (o *DocumentProcessorRunsDeleteRequest) GetRunID() string {
	if o == nil {
		return ""
	}
	return o.RunID
}

type DocumentProcessorRunsDeleteResponse struct {
	HTTPMeta                           components.HTTPMetadata `json:"-"`
	DeleteDocumentProcessorRunResponse *components.DeleteDocumentProcessorRunResponse
}

func (o *DocumentProcessorRunsDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DocumentProcessorRunsDeleteResponse) GetDeleteDocumentProcessorRunResponse() *components.DeleteDocumentProcessorRunResponse {
	if o == nil {
		return nil
	}
	return o.DeleteDocumentProcessorRunResponse
}