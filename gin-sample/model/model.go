package model

import "time"

type Interface interface {
	TableName() string
}

type UserInfo struct {
	ID          uint32    `gorm:"column:id;primaryKey" json:"id"`
	UserId      string    `gorm:"column:user_id;size:16;not null" json:"userId"`
	UserName    string    `gorm:"column:username;size:48;not null" json:"username"`
	NickName    string    `gorm:"column:nickname;size:48;not null" json:"nickname"`
	Password    string    `gorm:"column:password;size:16;not null" json:"password"`
	Email       string    `gorm:"column:email;size:48" json:"email"`
	CreatedTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdatedTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted   bool      `gorm:"column:is_deleted" json:"isDeleted"`
}

func (u *UserInfo) TableName() string {
	return TableUserInfo
}
