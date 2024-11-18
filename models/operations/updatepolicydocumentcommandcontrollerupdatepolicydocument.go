// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type UpdatePolicyDocumentCommandControllerUpdatePolicyDocumentRequest struct {
	// Unique identifier of the document.
	DocumentID string `pathParam:"style=simple,explode=false,name=documentId"`
	// Unique identifier of the policy.
	PolicyID                           string                                        `pathParam:"style=simple,explode=false,name=policyId"`
	UpdatePolicyDocumentCommandRequest components.UpdatePolicyDocumentCommandRequest `request:"mediaType=application/json"`
}

func (o *UpdatePolicyDocumentCommandControllerUpdatePolicyDocumentRequest) GetDocumentID() string {
	if o == nil {
		return ""
	}
	return o.DocumentID
}

func (o *UpdatePolicyDocumentCommandControllerUpdatePolicyDocumentRequest) GetPolicyID() string {
	if o == nil {
		return ""
	}
	return o.PolicyID
}

func (o *UpdatePolicyDocumentCommandControllerUpdatePolicyDocumentRequest) GetUpdatePolicyDocumentCommandRequest() components.UpdatePolicyDocumentCommandRequest {
	if o == nil {
		return components.UpdatePolicyDocumentCommandRequest{}
	}
	return o.UpdatePolicyDocumentCommandRequest
}

type UpdatePolicyDocumentCommandControllerUpdatePolicyDocumentResponse struct {
	HTTPMeta                            components.HTTPMetadata `json:"-"`
	UpdatePolicyDocumentCommandResponse *components.UpdatePolicyDocumentCommandResponse
}

func (o *UpdatePolicyDocumentCommandControllerUpdatePolicyDocumentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdatePolicyDocumentCommandControllerUpdatePolicyDocumentResponse) GetUpdatePolicyDocumentCommandResponse() *components.UpdatePolicyDocumentCommandResponse {
	if o == nil {
		return nil
	}
	return o.UpdatePolicyDocumentCommandResponse
}
