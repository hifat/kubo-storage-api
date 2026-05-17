package storagemdl

type UploadRequest struct {
	ObjectKey   string `json:"-"`
	Body        []byte `json:"body"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
}
