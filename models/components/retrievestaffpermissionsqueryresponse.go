// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RetrieveStaffPermissionsQueryResponse struct {
	// List of all staff roles
	Items []RoleAssignmentAggregate `json:"items"`
}

func (o *RetrieveStaffPermissionsQueryResponse) GetItems() []RoleAssignmentAggregate {
	if o == nil {
		return []RoleAssignmentAggregate{}
	}
	return o.Items
}