package request

type RegisterRequest struct {
	ClientID string `json:"client_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
