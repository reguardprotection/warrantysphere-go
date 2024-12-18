// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdatePolicyDocumentCommandResponse struct {
	// Unique identifier of the document.
	ID string `json:"id"`
	// Url at which the document is stored.
	UploadURL string `json:"uploadUrl"`
	// Unique identifier of the policy the coverage belongs to
	PolicyID string `json:"policyId"`
}

func (o *UpdatePolicyDocumentCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdatePolicyDocumentCommandResponse) GetUploadURL() string {
	if o == nil {
		return ""
	}
	return o.UploadURL
}

func (o *UpdatePolicyDocumentCommandResponse) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}
