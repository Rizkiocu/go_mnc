package dto

type AuthRequest struct {
	Email    string
	Password string
	Name     string
}

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
