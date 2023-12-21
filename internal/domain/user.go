package domain

//go:generate reform

//reform:users
type User struct {
	Id       int    `json:"Id,omitempty" reform:"id,pk"`
	Email    string `json:"Email" reform:"email"`
	Username string `json:"Username" reform:"username"`
	Password string `json:"Password" reform:"password"`
}
