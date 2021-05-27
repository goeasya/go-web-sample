package impl

import (
	"time"

	"gin-sample/model"
	"gin-sample/types"
	"gin-sample/util"

	"gorm.io/gorm"
)

type UserInfoImpl struct {
	DB *gorm.DB
}

func (u *UserInfoImpl) List() ([]*model.UserInfo, error) {
	var userList []*model.UserInfo
	err := u.DB.Where("is_deleted = ?", false).Find(&userList).Error
	return userList, err
}

func (u *UserInfoImpl) Create(user *model.UserInfo) error {
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	user.IsDeleted = false
	user.UserId = util.GenerateRandomID(16, util.RandomKindUpperLower)

	return u.DB.Create(&user).Error
}

func (u *UserInfoImpl) Update(userId string, user *model.UserInfo) error {
	user.UpdatedTime = time.Now()
	return u.DB.Table(model.TableUserInfo).Where("user_id = ?", userId).Updates(user).Error
}

func (u *UserInfoImpl) Delete(userId string) error {
	return u.DB.Table(model.TableUserInfo).
		Where("user_id = ?", userId).
		Update("is_deleted", true).
		Error
}

func (u *UserInfoImpl) GetByUserName(userName string) (*model.UserInfo, error) {
	var user model.UserInfo
	err := u.DB.Where("username = ? AND is_deleted =?", userName, false).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, types.ErrUserNotFound
	}
	return &user, err
}

func (u *UserInfoImpl) GetByUserId(userId string) (*model.UserInfo, error) {
	var user model.UserInfo
	err := u.DB.Where("user_id = ? AND is_deleted = ?", userId, false).First(&user).Error
	return &user, err
}
