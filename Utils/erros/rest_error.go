package erros

import "net/http"

// RestErr ...
type RestErr struct {
	Message string
	Status  int
	Error   string
}

//InternalServerError ...
func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

}

// BadRequestError ...
func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}

}
