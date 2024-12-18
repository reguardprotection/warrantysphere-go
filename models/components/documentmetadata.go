// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DocumentMetadata struct {
	// The original name of the file
	Filename string `json:"filename"`
	// The type of file
	Extension string `json:"extension"`
	// The file size
	Size float64 `json:"size"`
	// The expiry date of the download url, converted to number
	Expires *float64 `json:"expires"`
	// The name of the GCP storage bucket
	BucketName string `json:"bucketName"`
	// The access level of the GCP storage bucket
	IsPublic bool `json:"isPublic"`
	// The path of the stored file in GCP
	StoragePath string `json:"storagePath"`
}

func (o *DocumentMetadata) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

func (o *DocumentMetadata) GetExtension() string {
	if o == nil {
		return ""
	}
	return o.Extension
}

func (o *DocumentMetadata) GetSize() float64 {
	if o == nil {
		return 0.0
	}
	return o.Size
}

func (o *DocumentMetadata) GetExpires() *float64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *DocumentMetadata) GetBucketName() string {
	if o == nil {
		return ""
	}
	return o.BucketName
}

func (o *DocumentMetadata) GetIsPublic() bool {
	if o == nil {
		return false
	}
	return o.IsPublic
}

func (o *DocumentMetadata) GetStoragePath() string {
	if o == nil {
		return ""
	}
	return o.StoragePath
}
