package models

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=1,max=50"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserResponse struct {
	ResponseStatus ResponseStatus `json:"status"`
}

type LoginResponse struct {
	Data           LoginResponseData `json:"data"`
	ResponseStatus ResponseStatus    `json:"status"`
}

type LoginResponseData struct {
	AccessToken string `json:"access_token"`
}
