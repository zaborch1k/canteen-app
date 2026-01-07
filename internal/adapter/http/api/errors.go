package api

type InternalServerErrorResponse struct {
	Error string `json:"error" example:"internal server error"`
}
type InvalidRequestErrorResponse struct {
	Error string `json:"error" example:"invalid request"`
}

type InvalidCredentialsErrorResponse struct {
	Error string `json:"error" example:"invalid credentials"`
}

type UserExistsErrorResponse struct {
	Error string `json:"error" example:"user already exists"`
}

type RefreshTokenErrorResponse struct {
	Error string `json:"error" example:"refresh token error"`
}
