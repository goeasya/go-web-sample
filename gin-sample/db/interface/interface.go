package _interface

import "gin-sample/model"

type Database interface {
	InitDataBase()

	Users() UserInterface
}

type UserInterface interface {
	List() ([]*model.UserInfo, error)
	Create(user *model.UserInfo) error
	Update(userId string, user *model.UserInfo) error
	Delete(userId string) error
	GetByUserName(userName string) (*model.UserInfo, error)
	GetByUserId(userId string) (*model.UserInfo, error)
}
