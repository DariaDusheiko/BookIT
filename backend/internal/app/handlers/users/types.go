package users

type AuthRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=50"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}