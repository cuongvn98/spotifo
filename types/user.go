package types

type User struct {
	Email     string `json:"email"`
	FullName  string `json:"fullname"`
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	Avatar    string `json:"avatar"`
}
