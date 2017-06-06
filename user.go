package client

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Active   bool   `json:"active"`
	Expiry   string `json:"expiry,omitempty"`
	Created  string `json:"created,omitempty"`
}
