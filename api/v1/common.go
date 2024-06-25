package api

import (
	"encoding/json"
	"errors"
	"re-mall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON格式不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
