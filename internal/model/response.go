package model

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponse(msg string, status int) Response {

	return Response{
		Message: msg,
		Status:  status,
	}
}
