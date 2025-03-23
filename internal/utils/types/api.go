package types

type ApiResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{
		Error:   false,
		Message: "",
		Data:    nil,
	}
}
