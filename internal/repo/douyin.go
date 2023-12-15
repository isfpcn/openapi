package repo

import (
	"context"
	"openapi/api/douyin"
	"openapi/internal/config"
	"openapi/internal/dao"
	"openapi/pkg/tools"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

type DouyinRepo struct {
	DouyinTrade
}

const (
	DOUYIN_API_HOST         = "https://developer.toutiao.com"
	DOUYIN_SANDBOX_API_HOST = "https://open-sandbox.douyin.com"
	DOUYIN_PRODUCT_API_HOST = "https://developer-product.zijieapi.com/"

	//抖音client token 全局域名
	DOUYIN_CLIENT_TOKEN_API = "https://open.douyin.com"
)

func (r *DouyinRepo) GetClientToken(ctx context.Context, request *douyin.GetClientTokenRequest) (*douyin.GetClientTokenResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.ClientKey)
	if err != nil {
		return nil, err
	}
	//缓存的key
	key := "douyin:client_token:" + request.ClientKey
	ttl, err := dao.Redis.Do("TTL", key)
	if err != nil {
		return nil, err
	}

	rep := &douyin.GetClientTokenResponse{}
	switch ttl.(type) {
	case int64:
		if ttl.(int64) > 0 {
			value, err2 := dao.Redis.Do("GET", key)
			if err2 != nil {
				return nil, err2
			}
			err = json.Unmarshal([]byte(value.(string)), rep)
			if err != nil {
				return nil, err
			}
			rep.Data.ExpiresIn = ttl.(int64)
			return rep, nil
		}
	}

	reqData := make(map[string]string)

	reqData["client_key"] = request.ClientKey
	reqData["client_secret"] = info.Secret
	reqData["grant_type"] = "client_credential"

	err = r.Request("/oauth/client_token/", reqData, info.IsSandbox, rep)
	if err != nil {
		return nil, err
	}
	//缓存
	if rep != nil && rep.Data.ErrorCode == 0 {
		marshal, err := json.Marshal(rep)
		if err != nil {
			return nil, err
		}
		//有效期缩短10分钟
		expiresIn := rep.Data.ExpiresIn - 600
		_, err = dao.Redis.Set(key, string(marshal), time.Second*time.Duration(expiresIn))
		if err != nil {
			return nil, err
		}
		rep.Data.ExpiresIn = expiresIn
	}
	return rep, nil
}

func (r *DouyinRepo) GetAccessToken(ctx context.Context, request *douyin.GetAccessTokenRequest) (*douyin.GetAccessTokenResponse, error) {

	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.GetAccessTokenResponse{}

	key := "douyin:access_token:" + request.Appid
	ttl, err := dao.Redis.Do("TTL", key)

	if err != nil {
		return nil, err
	}

	switch ttl.(type) {
	case int64:
		if ttl.(int64) > 0 {
			get, err := dao.Redis.Do("GET", key)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(get.(string)), out)
			out.Data.ExpiresIn = ttl.(int64)
			return out, err
		}
	}

	data := make(map[string]string)

	data["appid"] = request.Appid
	data["secret"] = info.Secret
	data["grant_type"] = request.GrantType

	err = r.Request("/api/apps/v2/token", data, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	//缓存
	if out.ErrNo == 0 {
		marshal, err := json.Marshal(out)

		if err != nil {
			return nil, err
		}
		//有效期缩短10分钟
		expiresIn := out.Data.ExpiresIn - 600
		_, err = dao.Redis.Set(key, string(marshal), time.Second*time.Duration(expiresIn))
		if err != nil {
			return nil, err
		}
		out.Data.ExpiresIn = expiresIn
	}

	return out, nil
}

func (r *DouyinRepo) Request(path string, data interface{}, IsSandbox bool, out interface{}) error {
	host := DOUYIN_API_HOST
	if strings.Contains(path, "product/api") {
		host = DOUYIN_PRODUCT_API_HOST
	}
	if strings.Contains(path, "oauth/client_token") {
		host = DOUYIN_CLIENT_TOKEN_API
	}
	if IsSandbox {
		host = DOUYIN_SANDBOX_API_HOST
	}
	url := host + path
	marshal, err := json.Marshal(data)

	if err != nil {
		return err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	//沙盒环境自动审核参数，REJECT代表审核拒绝，其他代表审核成功
	if IsSandbox {
		headers["Aweme-Check-Type"] = "OK"
	}

	request, err := tools.HttpRequest("POST", url, marshal, headers)
	//fmt.Println("request = ", string(request))
	err = json.Unmarshal(request, out)
	if err != nil {
		return err
	}

	return nil
}

func (r *DouyinRepo) UploadMaterial(ctx context.Context, request *douyin.UploadMaterialRequest) (*douyin.UploadMaterialResponse, error) {

	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	file, err := tools.UrlToFile(request.MaterialFile)
	if err != nil {
		return nil, err
	}

	host := DOUYIN_API_HOST
	if info.IsSandbox {
		host = DOUYIN_SANDBOX_API_HOST
	}
	maps := make(map[string]string)
	maps["access_token"] = request.AccessToken
	maps["appid"] = request.Appid
	maps["material_type"] = strconv.Itoa(int(request.MaterialType))
	uploadFile, err := tools.UploadFile(file, host+"/auth/entity/upload_material", "material_file", maps)

	resp := &douyin.UploadMaterialResponse{}
	err = json.Unmarshal(uploadFile, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (r *DouyinRepo) BySelf(ctx context.Context, request *douyin.BySelfRequest) (*douyin.BySelfResponse, error) {

	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.BySelfResponse{}
	err = r.Request("/auth/entity/byself", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) AddRole(ctx context.Context, request *douyin.AddRoleRequest) (*douyin.AddRoleResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.AddRoleResponse{}
	err = r.Request("/auth/entity/add_role", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) Bypartner(ctx context.Context, request *douyin.BypartnerRequest) (*douyin.BypartnerResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.BypartnerResponse{}
	err = r.Request("/auth/entity/bypartner", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) GetBasicAuth(ctx context.Context, request *douyin.GetBasicAuthRequest) (*douyin.GetBasicAuthResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.GetBasicAuthResponse{}
	err = r.Request("/auth/entity/get_basic_auth", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) UpdateBasicAuth(ctx context.Context, request *douyin.UpdateBasicAuthRequest) (*douyin.UpdateBasicAuthResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.UpdateBasicAuthResponse{}
	err = r.Request("/auth/entity/update_basic_auth", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) GetClassAuth(ctx context.Context, request *douyin.GetClassAuthRequest) (*douyin.GetClassAuthResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.GetClassAuthResponse{}
	err = r.Request("/auth/entity/get_class_auth", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) UpdateClassAuth(ctx context.Context, request *douyin.UpdateClassAuthRequest) (*douyin.UpdateClassAuthResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.UpdateClassAuthResponse{}
	err = r.Request("/auth/entity/update_class_auth", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) AddClassAuth(ctx context.Context, request *douyin.AddClassAuthRequest) (*douyin.AddClassAuthResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.AddClassAuthResponse{}
	err = r.Request("/auth/entity/add_class_auth", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) GetAuditDetail(ctx context.Context, request *douyin.GetAuditDetailRequest) (*douyin.GetAuditDetailResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.GetAuditDetailResponse{}
	err = r.Request("/auth/entity/get_audit_detail", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) BindRole(ctx context.Context, request *douyin.BindRoleRequest) (*douyin.BindRoleResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.BindRoleResponse{}
	err = r.Request("/auth/entity/bind_role", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) UnbindRole(ctx context.Context, request *douyin.UnbindRoleRequest) (*douyin.UnbindRoleResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.UnbindRoleResponse{}
	err = r.Request("/auth/entity/un_bind_role", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) GetBindList(ctx context.Context, request *douyin.GetBindListRequest) (*douyin.GetBindListResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.GetBindListResponse{}
	err = r.Request("/auth/entity/get_bind_list", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) AuthRole(ctx context.Context, request *douyin.AuthRoleRequest) (*douyin.AuthRoleResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.AuthRoleResponse{}
	err = r.Request("/auth/entity/auth_role", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) UnauthRole(ctx context.Context, request *douyin.UnauthRoleRequest) (*douyin.UnauthRoleResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.UnauthRoleResponse{}
	err = r.Request("/auth/entity/un_auth_role", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) GetAppidAuth(ctx context.Context, request *douyin.GetAppidAuthRequest) (*douyin.GetAppidAuthResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.GetAppidAuthResponse{}
	err = r.Request("/auth/entity/get_appid_auth", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) UpdateAuthLetter(ctx context.Context, request *douyin.UpdateAuthLetterRequest) (*douyin.UpdateAuthLetterResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.UpdateAuthLetterResponse{}
	err = r.Request("/auth/entity/update_auth_letter", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) QueryMountscope(ctx context.Context, request *douyin.QueryMountscopeRequest) (*douyin.QueryMountscopeResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.QueryMountscopeResponse{}
	err = r.Request("/auth/entity/query_mount_scope", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) EnableMountscope(ctx context.Context, request *douyin.EnableMountscopeRequest) (*douyin.EnableMountscopeResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.EnableMountscopeResponse{}
	err = r.Request("/auth/entity/enable_mount_scope", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) UnbindAccount(ctx context.Context, request *douyin.UnbindAccountRequest) (*douyin.UnbindAccountResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.UnbindAccountResponse{}
	err = r.Request("/auth/entity/unbind_account", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) QueryEntityInfo(ctx context.Context, request *douyin.QueryEntityInfoRequest) (*douyin.QueryEntityInfoResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.AppId)
	if err != nil {
		return nil, err
	}

	out := &douyin.QueryEntityInfoResponse{}
	err = r.Request("/auth/entity/query_entity_info", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) QueryBindTocList(ctx context.Context, request *douyin.QueryBindTocListRequest) (*douyin.QueryBindTocListResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.QueryBindTocListResponse{}
	err = r.Request("/auth/entity/query_bind_toc_list", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) UploadResource(ctx context.Context, request *douyin.UploadResourceRequest) (*douyin.UploadResourceResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.UploadResourceResponse{}
	err = r.Request("/product/api/upload_resource", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) QueryResourceStatus(ctx context.Context, request *douyin.QueryResourceStatusRequest) (*douyin.QueryResourceStatusResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.QueryResourceStatusResponse{}
	err = r.Request("/product/api/query_resource_status", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) Add(ctx context.Context, request *douyin.AddRequest) (*douyin.AddResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.CommonProductParams.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.AddResponse{}
	err = r.Request("/product/api/add", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) Edit(ctx context.Context, request *douyin.EditRequest) (*douyin.EditResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Product.CommonProductParams.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.EditResponse{}
	err = r.Request("/product/api/edit", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) Query(ctx context.Context, request *douyin.QueryRequest) (*douyin.QueryResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.QueryResponse{}
	err = r.Request("/product/api/query", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) ModifyNoAudit(ctx context.Context, request *douyin.ModifyNoAuditRequest) (*douyin.ModifyNoAuditResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Product.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.ModifyNoAuditResponse{}
	err = r.Request("/product/api/modify_no_audit", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) QueryTemplateInfo(ctx context.Context, request *douyin.QueryTemplateInfoRequest) (*douyin.QueryTemplateInfoResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.QueryTemplateInfoResponse{}
	err = r.Request("/product/api/query_template_info", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) ModifyRefundRule(ctx context.Context, request *douyin.ModifyRefundRuleRequest) (*douyin.ModifyRefundRuleResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.ModifyRefundRuleResponse{}
	err = r.Request("/product/api/modify_refund_rule", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) QueryRefundRuleMeta(ctx context.Context, request *douyin.QueryRefundRuleMetaRequest) (*douyin.QueryRefundRuleMetaResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.QueryRefundRuleMetaResponse{}
	err = r.Request("/product/api/query_refund_rule_meta", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) QueryClassInfo(ctx context.Context, request *douyin.QueryClassInfoRequest) (*douyin.QueryClassInfoResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.QueryClassInfoResponse{}
	err = r.Request("/product/api/query_class_info", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *DouyinRepo) ModifyStatus(ctx context.Context, request *douyin.ModifyStatusRequest) (*douyin.ModifyStatusResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.ProductStatusInfo.Appid)
	if err != nil {
		return nil, err
	}

	out := &douyin.ModifyStatusResponse{}
	err = r.Request("/product/api/modify_status", request, info.IsSandbox, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}
