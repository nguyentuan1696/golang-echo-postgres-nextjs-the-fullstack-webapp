package response

import "thichlab-backend-docs/gerror"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	ErrorCode uint32 `json:"error_code"`
	Message   string `json:"message"`
	Exception string `json:"exception"`
}

func NewErrorResponse(errorCode uint32, message string, exception string) (string, ErrorResponse) {
	return gerror.StatusText(errorCode), ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
		Exception: exception,
	}
}
