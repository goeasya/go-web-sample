package handler

import (
	"time"

	"gin-sample/db"
	"gin-sample/global"
	"gin-sample/model"
	"gin-sample/model/request"
	"gin-sample/model/response"
	"gin-sample/types"
	"gin-sample/util"
)

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) RegisterUser(args *request.UserRegisterReq) error {
	var user model.UserInfo
	user.UserName = args.UserName
	user.NickName = args.NickName
	user.Password = util.EnCrypt(args.Password)
	user.Email = args.Email

	err := db.Manager().Users().Create(&user)
	if err != nil {
		global.Logger.Error("register user error: ", err.Error())
		return err
	}
	return nil
}

func (h *userHandler) Login(userName, password string) (*response.UserBase, error) {
	user, err := db.Manager().Users().GetByUserName(userName)
	if err != nil {
		global.Logger.Error("get user by username error: ", err.Error())
		return nil, err
	}
	if ok := util.ValidatePassword(password, user.Password); !ok {
		global.Logger.Error("user: %s, err: %s", userName, types.ErrInvalidUserOrPassword)
		return nil, types.ErrInvalidUserOrPassword
	}

	return &response.UserBase{
		UserId:   user.UserId,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
	}, nil
}

func (h *userHandler) DeleteUser(userId string) error {
	if err := db.Manager().Users().Delete(userId); err != nil {
		global.Logger.Error("delete user error: ", err.Error())
		return err
	}

	return nil
}

func (h *userHandler) UpdateUser(args *request.UserUpdateReq) error {
	user := &model.UserInfo{
		NickName:    args.NickName,
		Email:       args.Email,
		UpdatedTime: time.Now(),
	}

	err := db.Manager().Users().Update(args.UserId, user)
	if err != nil {
		global.Logger.Error("update user error: ", err.Error())
		return err
	}
	return nil
}

func (h *userHandler) GetUserDetail(username string) (*response.UserDetailResp, error) {
	user, err := db.Manager().Users().GetByUserName(username)
	if err != nil {
		global.Logger.Error("get user detail error: ", err.Error())
		return nil, err
	}

	resp := &response.UserDetailResp{
		UserId:     user.UserId,
		UserName:   user.UserName,
		NickName:   user.NickName,
		Email:      user.Email,
		CreateTime: user.CreatedTime.Format("2006-01-02 15:04:05"),
		UpdateTime: user.UpdatedTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}
