package storagehdl

type Handler struct {
	GRPC *GRPC
}

func New(GRPC *GRPC) *Handler {
	return &Handler{
		GRPC,
	}
}
