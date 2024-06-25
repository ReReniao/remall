package serializer

import (
	"re-mall/conf"
	"re-mall/model"
)

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	NikeName  string `json:"nike_name"`
	Type      int    `json:"type"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"createdAt"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:        user.ID,
		UserName:  user.UserName,
		NikeName:  user.NickName,
		Email:     user.Email,
		Status:    user.Status,
		Avatar:    conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
