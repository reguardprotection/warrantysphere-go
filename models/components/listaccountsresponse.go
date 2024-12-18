// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListAccountsResponse struct {
	Meta  PaginationMeta     `json:"meta"`
	Items []AccountAggregate `json:"items"`
}

func (o *ListAccountsResponse) GetMeta() PaginationMeta {
	if o == nil {
		return PaginationMeta{}
	}
	return o.Meta
}

func (o *ListAccountsResponse) GetItems() []AccountAggregate {
	if o == nil {
		return []AccountAggregate{}
	}
	return o.Items
}
