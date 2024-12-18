// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RetrievePortalConfigResponseConfig - The updated config for the organization
type RetrievePortalConfigResponseConfig struct {
}

type RetrievePortalConfigResponse struct {
	// The updated config for the organization
	Config RetrievePortalConfigResponseConfig `json:"config"`
	// Tenant logo.
	LogoUploadURL string `json:"logoUploadUrl"`
	// Tenant avatar.
	AvatarUploadURL string `json:"avatarUploadUrl"`
}

func (o *RetrievePortalConfigResponse) GetConfig() RetrievePortalConfigResponseConfig {
	if o == nil {
		return RetrievePortalConfigResponseConfig{}
	}
	return o.Config
}

func (o *RetrievePortalConfigResponse) GetLogoUploadURL() string {
	if o == nil {
		return ""
	}
	return o.LogoUploadURL
}

func (o *RetrievePortalConfigResponse) GetAvatarUploadURL() string {
	if o == nil {
		return ""
	}
	return o.AvatarUploadURL
}
