// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListWarrantiesResponseBody struct {
	Meta  PaginationMeta      `json:"meta"`
	Items []WarrantyAggregate `json:"items"`
}

func (o *ListWarrantiesResponseBody) GetMeta() PaginationMeta {
	if o == nil {
		return PaginationMeta{}
	}
	return o.Meta
}

func (o *ListWarrantiesResponseBody) GetItems() []WarrantyAggregate {
	if o == nil {
		return []WarrantyAggregate{}
	}
	return o.Items
}
