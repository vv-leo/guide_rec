package service

import (
	"errors"
	"time"
	"web3/constants"
	"web3/model"
)

var userDao = model.NewUserDao()

type UserService interface {
	Create(user *model.User) (success bool, err error)
	Detail(id string) (user *model.User, err error)
	Follow(from string, to string, status string) (success bool, err error)
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct {
}

func (s *userService) Create(user *model.User) (success bool, err error) {
	user.FansCount = 0
	user.CreateTime = time.Now().Unix()
	success, err = userDao.Create(user)
	return
}

func (s *userService) Detail(id string) (user *model.User, err error) {
	userDetail, err := userDao.DetailUser(id)
	return userDetail, err
}

func (s *userService) Follow(from string, to string, status string) (success bool, err error) {

	k := "follow:" + from + ":" + to
	v := model.GetRedisValue(k)

	if status == constants.FOLLOW {
		if len(v) != 0 {
			success, err = false, errors.New("already followed")
			return
		}
		model.SetRedisValue(k, "done", 0)
	} else {
		if len(v) == 0 {
			success, err = false, errors.New("already unfollowed")
			return
		}
		model.DelRedisKey(k)
	}
	success, err = userDao.FansCount(to, status)
	return
}
