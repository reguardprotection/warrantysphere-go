// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListPolicyPlansQueryResponse struct {
	Meta  PaginationMeta        `json:"meta"`
	Items []PolicyPlanAggregate `json:"items"`
}

func (o *ListPolicyPlansQueryResponse) GetMeta() PaginationMeta {
	if o == nil {
		return PaginationMeta{}
	}
	return o.Meta
}

func (o *ListPolicyPlansQueryResponse) GetItems() []PolicyPlanAggregate {
	if o == nil {
		return []PolicyPlanAggregate{}
	}
	return o.Items
}
