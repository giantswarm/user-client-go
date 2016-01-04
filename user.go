package client

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Expired  bool   `json:"expired,omitempty"`
}
