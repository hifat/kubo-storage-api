package storagemdl

type UploadRequest struct {
	ObjectKey string `json:"-"`
	Body      []byte `json:"body"`
}
