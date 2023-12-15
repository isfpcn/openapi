package douyin

import (
	"context"
	"openapi/api/douyin"
	"openapi/internal/repo"
)

var (
	_ douyin.DouyinTradeServer = (*TradeService)(nil)
)

type TradeService struct {
	Repo *repo.DouyinRepo
}

func (a *TradeService) MerchantAuditCallback(ctx context.Context, request *douyin.MerchantAuditCallbackRequest) (*douyin.MerchantAuditCallbackResponse, error) {
	//TODO implement me
	resp, err := a.Repo.MerchantAuditCallback(ctx, request)
	return resp, err
}

func (a *TradeService) CreateRefund(ctx context.Context, request *douyin.CreateRefundRequest) (*douyin.CreateREfundResponse, error) {
	//TODO implement me
	resp, err := a.Repo.CreateRefund(ctx, request)
	return resp, err
}

func (a *TradeService) QueryRefund(ctx context.Context, request *douyin.QueryRefundRequest) (*douyin.QueryRefundResponse, error) {
	//TODO implement me
	resp, err := a.Repo.QueryRefund(ctx, request)
	return resp, err
}

func (a *TradeService) QuerySettle(ctx context.Context, request *douyin.QuerySettleRequest) (*douyin.QuerySettleResponse, error) {
	//TODO implement me
	resp, err := a.Repo.QuerySettle(ctx, request)
	return resp, err
}

func (a *TradeService) QueryOrder(ctx context.Context, request *douyin.QueryOrderRequest) (*douyin.QueryOrderResponse, error) {
	//TODO implement me
	resp, err := a.Repo.QueryOrder(ctx, request)
	return resp, err
}

func (a *TradeService) QueryCps(ctx context.Context, request *douyin.QueryCpsRequest) (*douyin.QueryCpsResponse, error) {
	//TODO implement me
	resp, err := a.Repo.QueryCps(ctx, request)
	return resp, err
}
