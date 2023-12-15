package douyin

import (
	"context"
	"openapi/api/douyin"
	"openapi/internal/repo"
)

var (
	_ douyin.DouyinProductServer = (*ProductService)(nil)
)

type ProductService struct {
	Repo *repo.DouyinRepo
}

func (p *ProductService) UploadResource(ctx context.Context, request *douyin.UploadResourceRequest) (*douyin.UploadResourceResponse, error) {
	//TODO implement me
	resp, err := p.Repo.UploadResource(ctx, request)
	return resp, err
}

func (p *ProductService) QueryResourceStatus(ctx context.Context, request *douyin.QueryResourceStatusRequest) (*douyin.QueryResourceStatusResponse, error) {
	//TODO implement me
	resp, err := p.Repo.QueryResourceStatus(ctx, request)
	return resp, err
}

func (p *ProductService) Add(ctx context.Context, request *douyin.AddRequest) (*douyin.AddResponse, error) {
	//TODO implement me
	resp, err := p.Repo.Add(ctx, request)
	return resp, err
}

func (p *ProductService) Edit(ctx context.Context, request *douyin.EditRequest) (*douyin.EditResponse, error) {
	//TODO implement me
	resp, err := p.Repo.Edit(ctx, request)
	return resp, err
}

func (p *ProductService) Query(ctx context.Context, request *douyin.QueryRequest) (*douyin.QueryResponse, error) {
	//TODO implement me
	resp, err := p.Repo.Query(ctx, request)
	return resp, err
}

func (p *ProductService) ModifyNoAudit(ctx context.Context, request *douyin.ModifyNoAuditRequest) (*douyin.ModifyNoAuditResponse, error) {
	//TODO implement me
	resp, err := p.Repo.ModifyNoAudit(ctx, request)
	return resp, err
}

func (p *ProductService) QueryTemplateInfo(ctx context.Context, request *douyin.QueryTemplateInfoRequest) (*douyin.QueryTemplateInfoResponse, error) {
	//TODO implement me
	resp, err := p.Repo.QueryTemplateInfo(ctx, request)
	return resp, err
}

func (p *ProductService) ModifyRefundRule(ctx context.Context, request *douyin.ModifyRefundRuleRequest) (*douyin.ModifyRefundRuleResponse, error) {
	//TODO implement me
	resp, err := p.Repo.ModifyRefundRule(ctx, request)
	return resp, err
}

func (p *ProductService) QueryRefundRuleMeta(ctx context.Context, request *douyin.QueryRefundRuleMetaRequest) (*douyin.QueryRefundRuleMetaResponse, error) {
	//TODO implement me
	resp, err := p.Repo.QueryRefundRuleMeta(ctx, request)
	return resp, err
}

func (p *ProductService) QueryClassInfo(ctx context.Context, request *douyin.QueryClassInfoRequest) (*douyin.QueryClassInfoResponse, error) {
	//TODO implement me
	resp, err := p.Repo.QueryClassInfo(ctx, request)
	return resp, err
}

func (p *ProductService) ModifyStatus(ctx context.Context, request *douyin.ModifyStatusRequest) (*douyin.ModifyStatusResponse, error) {
	//TODO implement me
	resp, err := p.Repo.ModifyStatus(ctx, request)
	return resp, err
}
