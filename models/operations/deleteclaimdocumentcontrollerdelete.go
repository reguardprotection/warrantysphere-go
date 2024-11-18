// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type DeleteClaimDocumentControllerDeleteRequest struct {
	// The unique identifier of the claim.
	ClaimID string `pathParam:"style=simple,explode=false,name=claimId"`
	// Unique identifier of the document.
	DocumentID string `pathParam:"style=simple,explode=false,name=documentId"`
}

func (o *DeleteClaimDocumentControllerDeleteRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

func (o *DeleteClaimDocumentControllerDeleteRequest) GetDocumentID() string {
	if o == nil {
		return ""
	}
	return o.DocumentID
}

type DeleteClaimDocumentControllerDeleteResponse struct {
	HTTPMeta                      components.HTTPMetadata `json:"-"`
	DeleteDocumentCommandResponse *components.DeleteDocumentCommandResponse
}

func (o *DeleteClaimDocumentControllerDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteClaimDocumentControllerDeleteResponse) GetDeleteDocumentCommandResponse() *components.DeleteDocumentCommandResponse {
	if o == nil {
		return nil
	}
	return o.DeleteDocumentCommandResponse
}
