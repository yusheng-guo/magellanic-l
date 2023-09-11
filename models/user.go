package models

type User struct {
	ID       string `json:"id"`
	NickName string `json:"nick_name"` // 昵称
	Name     string `json:"name"`      // 姓名
	UserName string `json:"user_name"` // 用户名
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	IP       string `json:"ip"`
}

func (u *User) GetID() string {
	return u.ID
}
