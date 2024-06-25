package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// user模块错误
	ErrorExistUser             = 30001
	ErrorFailEncryption        = 30002
	ErrorExistUserNotFound     = 30003
	ErrorNotCompare            = 30004
	ErrorAuthToken             = 30005
	ErrorParseToken            = 30006
	ErrorAuthCheckTokenTimeout = 30007
	ErrorUploadFail            = 30008
	ErrorSendEmail             = 30009

	// product模块错误
	ErrorProductImgUpload = 40001

	// 收藏夹错误
	ErrorFavoriteExist = 50001
)
