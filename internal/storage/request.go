package storagemdl

type UploadRequest struct {
	ObjectKey   string `json:"-"`
	Body        []byte `json:"body"`
	Filename    string `json:"filename"`
	Path        string `json:"path"`
	ContentType string `json:"content_type"`
}
