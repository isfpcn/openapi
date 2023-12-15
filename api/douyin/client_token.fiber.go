// Code generated by protoc-gen-fiber v0.1, DO NOT EDIT.
// source: api/douyin/client_token.proto

package douyin

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/gofiber/fiber/v2"
)

var (
	PathDouyinOauthClientToken = "/douyin/oauth/client_token"
)

// DouyinClientTokenFiberServer is the server API for DouyinClientToken service.
// 抖音-交易client token调用凭证
type DouyinClientTokenFiberServer interface {
	// client_token 用于不需要用户授权就可以调用的接口。client_token 的有效时间为 2 个小时，重复获取 client_token 后会使上次的 client_token 失效（但有 5 分钟的缓冲时间，连续多次获取 client_token 只会保留最新的两个 client_token）
	GetClientToken(ctx context.Context, req *GetClientTokenRequest) (resp *GetClientTokenResponse, err error)
}

var (
	DouyinClientTokenSvc       DouyinClientTokenFiberServer
	DouyinClientTokenWriter    func(c *fiber.Ctx, message proto.Message) error
	DouyinClientTokenValidater func(message proto.Message) error
)

func _DouyinClientToken_GetClientToken0_HTTP_Handler(c *fiber.Ctx) error {
	p := new(GetClientTokenRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	if err := DouyinClientTokenValidater(p); err != nil {
		return err
	}
	resp, err := DouyinClientTokenSvc.GetClientToken(c.UserContext(), p)
	if err != nil {
		return err
	}
	return DouyinClientTokenWriter(c, resp)
}

// RegisterDouyinClientTokenFiberServer Register the fiber route
func RegisterDouyinClientTokenFiberServer(e *fiber.App, server DouyinClientTokenFiberServer, w func(c *fiber.Ctx, message proto.Message) error, v func(message proto.Message) error) {
	DouyinClientTokenSvc = server
	DouyinClientTokenWriter = w
	DouyinClientTokenValidater = v
	e.Post("/douyin/oauth/client_token", _DouyinClientToken_GetClientToken0_HTTP_Handler)
}