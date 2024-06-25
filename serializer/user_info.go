package serializer

import (
	"context"
	"errors"
)

type Key int

var userKey Key

type UserInfo struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
}

func NewContext(c context.Context, user *UserInfo) context.Context {
	return context.WithValue(c, userKey, user)
}

func FromContext(c context.Context) (*UserInfo, bool) {
	value, ok := c.Value(userKey).(*UserInfo)
	return value, ok
}

func GetUserInfo(c context.Context) (*UserInfo, error) {
	userInfo, ok := FromContext(c)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return userInfo, nil
}
