package douyin

import (
	"context"
	"openapi/api/douyin"
	"openapi/internal/repo"
)

var (
	_ douyin.DouyinAccessTokenServer = (*AccessTokenService)(nil)
)

type AccessTokenService struct {
	Repo *repo.DouyinRepo
}

func (a *AccessTokenService) GetAccessToken(ctx context.Context, request *douyin.GetAccessTokenRequest) (*douyin.GetAccessTokenResponse, error) {
	//TODO implement me

	resp, err := a.Repo.GetAccessToken(ctx, request)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
