// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdateClaimNoteCommandResponse struct {
	// Unique identifier of the claim note.
	ID string `json:"id"`
}

func (o *UpdateClaimNoteCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
