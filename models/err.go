package models

import (
	"encoding/json"
	"fmt"
)

var (
	Err_invalid_param  = AppError{Code: "4001", Message: "Invalid parameters"}
	Err_required_token = AppError{Code: "4002", Message: "required service token"}

	Err_backend_system = AppError{Code: "5001", Message: "Error from backend system"}
	Err_data_not_found = AppError{Code: "5002", Message: "Error data not found"}
	Err_duplicate_data = AppError{Code: "5003", Message: "Error duplicate data"}
	Err_incorrect_data = AppError{Code: "5004", Message: "Error incorrect data"}
	Err_cannot_delete  = AppError{Code: "5005", Message: "Error cannot delete"}
	Err_cannot_vote    = AppError{Code: "5006", Message: "Error cannot vote"}

	Unknown = AppError{Code: "9999", Message: "Unknown error"}
)

type AppError struct {
	Code    string `json:"app_code"`
	Message string `json:"app_message"`
	Value   interface{}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("app error:%s|%s|%+v", e.Code, e.Message, e.Value)
}

func (i *AppError) String() string {
	str, _ := json.Marshal(i)
	return string(str)
}

func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

type ErrorResponse struct {
	Error ErrorWebStatus `json:"error"`
}

type ErrorWebStatus struct {
	Code    string `json:"app_code"`
	Message string `json:"app_message"`
}

func (e *ErrorResponse) String() string {
	str, _ := json.Marshal(e)
	return string(str)
}

func NewErrorWebResponse(err error) ErrorResponse {
	var code, message string
	switch v := err.(type) {
	case *AppError:
		code = v.Code
		message = v.Message
	default:
		code = Unknown.Code
		message = Unknown.Message
	}
	return ErrorResponse{
		Error: ErrorWebStatus{
			Code:    code,
			Message: message,
		},
	}
}
