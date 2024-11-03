package model

import (
	"errors"
	"gorm.io/gorm"
	"web3/constants"
)

type User struct {
	Id          string `json:"id"`
	FansCount   int    `json:"fans_count"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time"`
}

func (User) TableName() string {
	return "user"
}

type UserDao interface {
	Create(user *User) (success bool, err error)
	DetailUser(id string) (user *User, err error)
	FansCount(id string, status string) (success bool, err error)
}

func NewUserDao() UserDao {
	return &userDao{}
}

type userDao struct {
}

func (s *userDao) operateTable() *gorm.DB {
	return db.Table("user")
}

func (s *userDao) Create(user *User) (success bool, err error) {
	if res := s.operateTable().Create(user); res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (s *userDao) DetailUser(id string) (user *User, err error) {
	if res := s.operateTable().Where("id", id).Find(&user); res.Error != nil {
		// 检查错误类型，确定是否是记录未找到
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			// 返回一个空的用户对象
			return nil, nil // 返回空的用户实例
		}
		return nil, res.Error
	}
	if user.Id == "" || len(user.Id) == 0 {
		return nil, nil
	}
	return user, nil
}

func (s *userDao) FansCount(id string, status string) (success bool, err error) {
	var user User
	if resF := s.operateTable().First(&user, id); resF.Error != nil {
		return false, resF.Error
	}
	if constants.FOLLOW == status {
		user.FansCount++
	} else {
		user.FansCount--
	}
	if resS := db.Save(user); resS.Error != nil {
		return false, resS.Error
	}
	return true, nil
}
