package res

type SignIn struct {
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
	ClientID  string `json:"client_id"`
	Status    int16  `json:"status"`
	Sysadmin  int16  `json:"sysadmin"`
	UserToken string `json:"user_token"`
}
