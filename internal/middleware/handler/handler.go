package middlewarehdl

type Handler struct {
	GRPC *GRPC
}

func New(grpc *GRPC) *Handler {
	return &Handler{
		GRPC: grpc,
	}
}
