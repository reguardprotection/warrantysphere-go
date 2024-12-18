// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DeleteCoverageCommandRequest struct {
	// Whether or not the command should create a new draft policy if required.
	CanCreateNewVersion bool `json:"canCreateNewVersion"`
}

func (o *DeleteCoverageCommandRequest) GetCanCreateNewVersion() bool {
	if o == nil {
		return false
	}
	return o.CanCreateNewVersion
}
