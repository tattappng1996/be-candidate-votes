package models

var (
	SUCCESSResponse = ResponseStatus{Code: "0000", Message: "Success"}
)

type GeneralResponse struct {
	ResponseStatus ResponseStatus `json:"status"`
}

type ResponseStatus struct {
	Code    string `json:"app_code"`
	Message string `json:"app_message"`
}
