package service

import (
	"context"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"re-mall/conf"
	"re-mall/dao"
	"re-mall/model"
	"re-mall/pkg/e"
	"re-mall/pkg/util"
	"re-mall/serializer"
	"strings"
	"time"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` // 前端验证
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	//1 绑定邮箱 2 解绑邮箱 3 更改密码
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

// Register 注册
func (s *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if s.Key == "" || len(s.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度过短",
		}
	}
	// 10000 ---> 密文加密 对称加密操作
	util.Encrypt.SetKey(s.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(s.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: s.UserName,
		NickName: s.NickName,
		Status:   model.Active,
		Avatar:   "favicon.ico",
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额的加密
	}

	// 密码加密
	if err = user.SetPassword(s.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 登录
func (s *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(s.UserName)
	if !exist || err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	if user.CheckPassword(s.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	// token 的签发
	token, err := util.GenerateTaken(user.ID, s.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	return serializer.Response{
		Status: code,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg:   e.GetMsg(code),
		Error: "",
	}
}

// Update 用户修改信息
func (s *UserService) Update(ctx context.Context) serializer.Response {
	code := e.Success
	u, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(u.Id)
	// 修改用户昵称 nickname
	if s.NickName != user.NickName {
		user.NickName = s.NickName
	}
	err = userDao.UpdateUserById(u.Id, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}

// Post 上传头像
func (s *UserService) Post(ctx context.Context, file multipart.File, filesize int64) serializer.Response {
	code := e.Success
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(userInfo.Id)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	// 保存图片到本地
	path, err := UploadAvatarToLocalStatic(file, userInfo.Id, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(user.ID, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   nil,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}

// Send 发送邮件
func (s *SendEmailService) Send(ctx context.Context) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice // 邮件模版
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}

	token, err := util.GenerateEmailTaken(userInfo.Id, s.OperationType, s.Email, s.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(s.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	address = conf.ValidEmail + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", s.Email)
	m.SetHeader("Subject", "Reniao")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   nil,
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}

// Valid 验证邮箱
func (s *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.Success
	// 验证token
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorParseToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}

	// 获取该用户的信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	if operationType == 1 {
		// 绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		// 解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  "",
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}

// ShowMoney 展示用户金额
func (s *ShowMoneyService) Show(ctx context.Context) serializer.Response {
	code := e.Success
	userInfo, err := serializer.GetUserInfo(ctx)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userInfo.Id)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, s.Key),
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}
