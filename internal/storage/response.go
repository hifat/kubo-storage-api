package storagemdl

type UploadResponse struct {
	ObjectKey string `json:"object_key"`
	URL       string `json:"url"`
}

type PresignedResponse struct {
	URL string `json:"url"`
}
