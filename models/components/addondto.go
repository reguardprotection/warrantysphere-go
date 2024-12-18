// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AddonDto struct {
	// Unique identifier of the addon.
	ID string `json:"id"`
	// Title of the addon.
	Title string `json:"title"`
	// Description of the addon.
	Description string `json:"description"`
}

func (o *AddonDto) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AddonDto) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *AddonDto) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}
