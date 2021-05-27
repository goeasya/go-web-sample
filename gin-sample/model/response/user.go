package response

type UserBase struct {
	UserId   string `json:"userId"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
}

type LoginToken struct {
	Token string `json:"token"`
}

type UserDetailResp struct {
	UserId     string `json:"userId"`
	UserName   string `json:"username"`
	NickName   string `json:"nickname"`
	Email      string `json:"email"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
