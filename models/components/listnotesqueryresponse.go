// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ListNotesQueryResponse struct {
	Meta  PaginationMeta  `json:"meta"`
	Items []NoteAggregate `json:"items"`
}

func (o *ListNotesQueryResponse) GetMeta() PaginationMeta {
	if o == nil {
		return PaginationMeta{}
	}
	return o.Meta
}

func (o *ListNotesQueryResponse) GetItems() []NoteAggregate {
	if o == nil {
		return []NoteAggregate{}
	}
	return o.Items
}
