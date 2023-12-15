/**
 * @author   Liang
 * @create   2023/11/30 18:02
 * @version  1.0
 */

package douyin

import (
	"context"
	"openapi/api/douyin"
	"openapi/internal/repo"
)

var (
	_ douyin.DouyinAuthServer = (*AuthService)(nil)
)

type AuthService struct {
	Repo *repo.DouyinRepo
}

func (a *AuthService) GetBasicAuth(ctx context.Context, request *douyin.GetBasicAuthRequest) (*douyin.GetBasicAuthResponse, error) {
	//TODO implement me
	resp, err := a.Repo.GetBasicAuth(ctx, request)
	return resp, err
}

func (a *AuthService) UpdateBasicAuth(ctx context.Context, request *douyin.UpdateBasicAuthRequest) (*douyin.UpdateBasicAuthResponse, error) {
	//TODO implement me
	resp, err := a.Repo.UpdateBasicAuth(ctx, request)
	return resp, err
}

func (a *AuthService) GetClassAuth(ctx context.Context, request *douyin.GetClassAuthRequest) (*douyin.GetClassAuthResponse, error) {
	//TODO implement me
	resp, err := a.Repo.GetClassAuth(ctx, request)
	return resp, err
}

func (a *AuthService) UpdateClassAuth(ctx context.Context, request *douyin.UpdateClassAuthRequest) (*douyin.UpdateClassAuthResponse, error) {
	//TODO implement me
	resp, err := a.Repo.UpdateClassAuth(ctx, request)
	return resp, err
}

func (a *AuthService) AddClassAuth(ctx context.Context, request *douyin.AddClassAuthRequest) (*douyin.AddClassAuthResponse, error) {
	//TODO implement me
	resp, err := a.Repo.AddClassAuth(ctx, request)
	return resp, err
}

func (a *AuthService) GetAuditDetail(ctx context.Context, request *douyin.GetAuditDetailRequest) (*douyin.GetAuditDetailResponse, error) {
	//TODO implement me
	resp, err := a.Repo.GetAuditDetail(ctx, request)
	return resp, err
}

func (a *AuthService) BindRole(ctx context.Context, request *douyin.BindRoleRequest) (*douyin.BindRoleResponse, error) {
	//TODO implement me
	resp, err := a.Repo.BindRole(ctx, request)
	return resp, err
}

func (a *AuthService) UnbindRole(ctx context.Context, request *douyin.UnbindRoleRequest) (*douyin.UnbindRoleResponse, error) {
	//TODO implement me
	resp, err := a.Repo.UnbindRole(ctx, request)
	return resp, err
}

func (a *AuthService) GetBindList(ctx context.Context, request *douyin.GetBindListRequest) (*douyin.GetBindListResponse, error) {
	//TODO implement me
	resp, err := a.Repo.GetBindList(ctx, request)
	return resp, err
}

func (a *AuthService) AuthRole(ctx context.Context, request *douyin.AuthRoleRequest) (*douyin.AuthRoleResponse, error) {
	//TODO implement me
	resp, err := a.Repo.AuthRole(ctx, request)
	return resp, err
}

func (a *AuthService) UnauthRole(ctx context.Context, request *douyin.UnauthRoleRequest) (*douyin.UnauthRoleResponse, error) {
	//TODO implement me
	resp, err := a.Repo.UnauthRole(ctx, request)
	return resp, err
}

func (a *AuthService) GetAppidAuth(ctx context.Context, request *douyin.GetAppidAuthRequest) (*douyin.GetAppidAuthResponse, error) {
	//TODO implement me
	resp, err := a.Repo.GetAppidAuth(ctx, request)
	return resp, err
}

func (a *AuthService) UpdateAuthLetter(ctx context.Context, request *douyin.UpdateAuthLetterRequest) (*douyin.UpdateAuthLetterResponse, error) {
	//TODO implement me
	resp, err := a.Repo.UpdateAuthLetter(ctx, request)
	return resp, err
}

func (a *AuthService) QueryMountscope(ctx context.Context, request *douyin.QueryMountscopeRequest) (*douyin.QueryMountscopeResponse, error) {
	//TODO implement me
	resp, err := a.Repo.QueryMountscope(ctx, request)
	return resp, err
}

func (a *AuthService) EnableMountscope(ctx context.Context, request *douyin.EnableMountscopeRequest) (*douyin.EnableMountscopeResponse, error) {
	//TODO implement me
	resp, err := a.Repo.EnableMountscope(ctx, request)
	return resp, err
}

func (a *AuthService) UnbindAccount(ctx context.Context, request *douyin.UnbindAccountRequest) (*douyin.UnbindAccountResponse, error) {
	//TODO implement me
	resp, err := a.Repo.UnbindAccount(ctx, request)
	return resp, err
}

func (a *AuthService) QueryEntityInfo(ctx context.Context, request *douyin.QueryEntityInfoRequest) (*douyin.QueryEntityInfoResponse, error) {
	//TODO implement me
	resp, err := a.Repo.QueryEntityInfo(ctx, request)
	return resp, err
}

func (a *AuthService) QueryBindTocList(ctx context.Context, request *douyin.QueryBindTocListRequest) (*douyin.QueryBindTocListResponse, error) {
	//TODO implement me
	resp, err := a.Repo.QueryBindTocList(ctx, request)
	return resp, err
}

func (a *AuthService) BySelf(ctx context.Context, request *douyin.BySelfRequest) (*douyin.BySelfResponse, error) {
	//TODO implement me
	resp, err := a.Repo.BySelf(ctx, request)
	return resp, err
}

func (a *AuthService) AddRole(ctx context.Context, request *douyin.AddRoleRequest) (*douyin.AddRoleResponse, error) {
	//TODO implement me
	resp, err := a.Repo.AddRole(ctx, request)
	return resp, err
}

func (a *AuthService) Bypartner(ctx context.Context, request *douyin.BypartnerRequest) (*douyin.BypartnerResponse, error) {
	//TODO implement me
	resp, err := a.Repo.Bypartner(ctx, request)
	return resp, err
}

func (a *AuthService) UploadMaterial(ctx context.Context, request *douyin.UploadMaterialRequest) (*douyin.UploadMaterialResponse, error) {
	resp, err := a.Repo.UploadMaterial(ctx, request)
	return resp, err
}
