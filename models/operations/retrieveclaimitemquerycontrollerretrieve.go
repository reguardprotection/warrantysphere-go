// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type RetrieveClaimItemQueryControllerRetrieveRequest struct {
	// Unique identifier of the claim item.
	ItemID string `pathParam:"style=simple,explode=false,name=itemId"`
	// Unique identifier of the claim associated with the claim item.
	ClaimID string `pathParam:"style=simple,explode=false,name=claimId"`
}

func (o *RetrieveClaimItemQueryControllerRetrieveRequest) GetItemID() string {
	if o == nil {
		return ""
	}
	return o.ItemID
}

func (o *RetrieveClaimItemQueryControllerRetrieveRequest) GetClaimID() string {
	if o == nil {
		return ""
	}
	return o.ClaimID
}

type RetrieveClaimItemQueryControllerRetrieveResponse struct {
	HTTPMeta           components.HTTPMetadata `json:"-"`
	ClaimItemAggregate *components.ClaimItemAggregate
}

func (o *RetrieveClaimItemQueryControllerRetrieveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveClaimItemQueryControllerRetrieveResponse) GetClaimItemAggregate() *components.ClaimItemAggregate {
	if o == nil {
		return nil
	}
	return o.ClaimItemAggregate
}