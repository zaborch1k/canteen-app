package common

type RegisterRequest struct {
	Login    string `json:"login" binding:"required" validate:"required,max=50,min=2" example:"the_real_slim_shady"`
	Password string `json:"password" binding:"required" validate:"required,max=100,min=8" example:"password1234"`
	Name     string `json:"name" binding:"required" validate:"required,max=100,alpha" example:"Slim"`
	Surname  string `json:"surname" binding:"required" validate:"required,max=100,alpha" example:"Shady"`
	Role     string `json:"role" binding:"required" validate:"required,oneof=admin employee user" example:"admin"`
}

type LoginRequest struct {
	Login    string `json:"login" binding:"required" validate:"required,max=50" example:"the_real_slim_shady"`
	Password string `json:"password" binding:"required" validate:"required,max=100" example:"password1234"`
}
