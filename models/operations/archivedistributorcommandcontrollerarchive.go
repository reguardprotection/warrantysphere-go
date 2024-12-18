// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/reguardprotection/warrantysphere/models/components"
)

type ArchiveDistributorCommandControllerArchiveRequest struct {
	// Unique identifier of the distributor.
	DistributorID string `pathParam:"style=simple,explode=false,name=distributorId"`
}

func (o *ArchiveDistributorCommandControllerArchiveRequest) GetDistributorID() string {
	if o == nil {
		return ""
	}
	return o.DistributorID
}

type ArchiveDistributorCommandControllerArchiveResponse struct {
	HTTPMeta                          components.HTTPMetadata `json:"-"`
	ArchiveDistributorCommandResponse *components.ArchiveDistributorCommandResponse
}

func (o *ArchiveDistributorCommandControllerArchiveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ArchiveDistributorCommandControllerArchiveResponse) GetArchiveDistributorCommandResponse() *components.ArchiveDistributorCommandResponse {
	if o == nil {
		return nil
	}
	return o.ArchiveDistributorCommandResponse
}
