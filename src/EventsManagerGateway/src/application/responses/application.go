package response

type badRequest struct {
	Message string
}

func NewBadRequest(err error) *badRequest {
	return &badRequest{
		Message: err.Error(),
	}
}
