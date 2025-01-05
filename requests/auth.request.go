package requests

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"` //email
	Password string `json:"password" form:"password" binding:"required"` //password
	Name     string `json:"name" form:"name" binding:"required"`         //name
}
