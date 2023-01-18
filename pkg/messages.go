package pkg

type Response struct {
	Status bool   `json:"status"`
	Data   any    `json:"data"`
	Error  string `json:"error"`
}

func SuccessResponse(data any) Response {
	return Response{
		Status: true,
		Data:   data,
		Error:  "",
	}
}

func ErrorResponse(err error) Response {
	return Response{
		Status: false,
		Data:   "",
		Error:  err.Error(),
	}
}
