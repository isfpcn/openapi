/**
 * @author   Liang
 * @create   2023/11/30 10:12
 * @version  1.0
 */

package grpc

import (
	"net"
	"openapi/api/douyin"
	"openapi/internal/config"
	"openapi/internal/repo"
	douyin2 "openapi/internal/service/douyin"
	"openapi/internal/transport/grpc/md"

	"google.golang.org/grpc"
)

func New() error {
	lis, err := net.Listen("tcp", config.CFG.Server.GrpcAddr)

	if err != nil {
		return err
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(md.MiddlewareInterceptor),
	)

	douyinRepo := new(repo.DouyinRepo)

	douyin.RegisterDouyinTradeServer(s, &douyin2.TradeService{Repo: douyinRepo})
	douyin.RegisterDouyinClientTokenServer(s, &douyin2.ClientTokenService{Repo: douyinRepo})
	douyin.RegisterDouyinAuthServer(s, &douyin2.AuthService{Repo: douyinRepo})
	douyin.RegisterDouyinAccessTokenServer(s, &douyin2.AccessTokenService{Repo: douyinRepo})
	douyin.RegisterDouyinProductServer(s, &douyin2.ProductService{Repo: douyinRepo})

	err = s.Serve(lis)
	return err
}
