package request

type UserRegisterReq struct {
	UserName string `json:"username" binding:"required,gte=6"`
	Password string `json:"password" binding:"required,gte=8"`
	NickName string `json:"nickname" binding:"required,gte=6"`
	Email    string `json:"email"`
}

type UserLoginReq struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDelReq struct {
	UserId string `json:"userId" binding:"required"`
}

type UserUpdateReq struct {
	UserId   string `json:"userId" binding:"required"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
}
