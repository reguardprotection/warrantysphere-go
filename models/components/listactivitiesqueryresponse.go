// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListActivitiesQueryResponse struct {
	Meta  PaginationMeta      `json:"meta"`
	Items []ActivityAggregate `json:"items"`
}

func (o *ListActivitiesQueryResponse) GetMeta() PaginationMeta {
	if o == nil {
		return PaginationMeta{}
	}
	return o.Meta
}

func (o *ListActivitiesQueryResponse) GetItems() []ActivityAggregate {
	if o == nil {
		return []ActivityAggregate{}
	}
	return o.Items
}