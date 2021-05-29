package credentials

type Credentials struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
}
