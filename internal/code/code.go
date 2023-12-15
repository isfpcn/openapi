package code

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	OK = codes.Code(0) //成功

	// 自定义错误码
	InvalidArgument = codes.Code(100) //参数错误
)

func Error(code codes.Code, msg string) error {
	return status.Error(code, msg)
}

func Codec(code uint32, msg string) error {
	return status.Error(codes.Code(code), msg)
}
