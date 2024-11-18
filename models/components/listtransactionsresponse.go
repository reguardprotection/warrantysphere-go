// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListTransactionsResponse struct {
	Meta  PaginationMeta         `json:"meta"`
	Items []TransactionAggregate `json:"items"`
}

func (o *ListTransactionsResponse) GetMeta() PaginationMeta {
	if o == nil {
		return PaginationMeta{}
	}
	return o.Meta
}

func (o *ListTransactionsResponse) GetItems() []TransactionAggregate {
	if o == nil {
		return []TransactionAggregate{}
	}
	return o.Items
}