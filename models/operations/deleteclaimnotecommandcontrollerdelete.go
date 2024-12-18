// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type DeleteClaimNoteCommandControllerDeleteRequest struct {
	// Unique identifier of the claim note.
	NoteID string `pathParam:"style=simple,explode=false,name=noteId"`
	// Unique identifier of the claim associated with the claim note.
	ClaimID string `pathParam:"style=simple,explode=false,name=claimId"`
}

func (o *DeleteClaimNoteCommandControllerDeleteRequest) GetNoteID() string {
	if o == nil {
		return ""
	}
	return o.NoteID
}

func (o *DeleteClaimNoteCommandControllerDeleteRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

type DeleteClaimNoteCommandControllerDeleteResponse struct {
	HTTPMeta                       components.HTTPMetadata `json:"-"`
	DeleteClaimNoteCommandResponse *components.DeleteClaimNoteCommandResponse
}

func (o *DeleteClaimNoteCommandControllerDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteClaimNoteCommandControllerDeleteResponse) GetDeleteClaimNoteCommandResponse() *components.DeleteClaimNoteCommandResponse {
	if o == nil {
		return nil
	}
	return o.DeleteClaimNoteCommandResponse
}
