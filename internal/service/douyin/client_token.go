package douyin

import (
	"context"
	"openapi/api/douyin"
	"openapi/internal/repo"
)

var (
	_ douyin.DouyinClientTokenServer = (*ClientTokenService)(nil)
)

type ClientTokenService struct {
	Repo *repo.DouyinRepo
}

func (a *ClientTokenService) GetClientToken(ctx context.Context, request *douyin.GetClientTokenRequest) (*douyin.GetClientTokenResponse, error) {
	//TODO implement me
	resp, err := a.Repo.GetClientToken(ctx, request)
	return resp, err
}
