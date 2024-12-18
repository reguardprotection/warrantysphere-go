// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type RetrieveClaimNoteQueryControllerRetrieveRequest struct {
	// Unique identifier of the claim used to filter the claim notes.
	ClaimID string `pathParam:"style=simple,explode=false,name=claimId"`
	// Unique identifier of the claim used to filter the claim notes.
	NoteID string `pathParam:"style=simple,explode=false,name=noteId"`
}

func (o *RetrieveClaimNoteQueryControllerRetrieveRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

func (o *RetrieveClaimNoteQueryControllerRetrieveRequest) GetNoteID() string {
	if o == nil {
		return ""
	}
	return o.NoteID
}

type RetrieveClaimNoteQueryControllerRetrieveResponse struct {
	HTTPMeta      components.HTTPMetadata `json:"-"`
	NoteAggregate *components.NoteAggregate
}

func (o *RetrieveClaimNoteQueryControllerRetrieveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveClaimNoteQueryControllerRetrieveResponse) GetNoteAggregate() *components.NoteAggregate {
	if o == nil {
		return nil
	}
	return o.NoteAggregate
}
