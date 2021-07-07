package service

import (
	"github.com/xszyh/gotest/app/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}
func (u *UserService) GetUserInfo(userId string) *model.UserInfo {
	return model.GetUserFromDb(userId)
}
