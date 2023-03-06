package model

// Model credentiales of the user request
type Credentiales struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
