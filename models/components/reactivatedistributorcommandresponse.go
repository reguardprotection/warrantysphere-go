// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ReactivateDistributorCommandResponseStatus - Status of the distributor after being reactivated.
type ReactivateDistributorCommandResponseStatus struct {
}

type ReactivateDistributorCommandResponse struct {
	// Unique identifier of the distributor.
	ID string `json:"id"`
	// Status of the distributor after being reactivated.
	Status ReactivateDistributorCommandResponseStatus `json:"status"`
}

func (o *ReactivateDistributorCommandResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ReactivateDistributorCommandResponse) GetStatus() ReactivateDistributorCommandResponseStatus {
	if o == nil {
		return ReactivateDistributorCommandResponseStatus{}
	}
	return o.Status
}