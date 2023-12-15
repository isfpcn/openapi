package repo

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"openapi/api/douyin"
	"openapi/internal/config"
	"openapi/pkg/tools"
	"strconv"
	"time"
)

type DouyinTrade struct {
}

func (r *DouyinTrade) MerchantAuditCallback(ctx context.Context, request *douyin.MerchantAuditCallbackRequest) (*douyin.MerchantAuditCallbackResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}
	body := &douyin.MerchantAuditCallbackRequest{
		OutRefundNo:       request.OutRefundNo,
		RefundAuditStatus: request.RefundAuditStatus,
		DenyMessage:       request.DenyMessage,
	}
	marshal, _ := json.Marshal(body)
	path := "/api/apps/trade/v2/merchant_audit_callback"
	res, err := r.sign(string(marshal), info.ApplicationPrivateKey, path)
	if err != nil {
		return nil, err
	}
	authorization := r.getAuthorization(request.Appid, "1", res)
	out := &douyin.MerchantAuditCallbackResponse{}
	_ = r.TradeRequest(path, marshal, info.IsSandbox, out, authorization)
	return out, nil
}

func (r *DouyinTrade) CreateRefund(ctx context.Context, request *douyin.CreateRefundRequest) (*douyin.CreateREfundResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}
	body := &douyin.CreateRefundRequest{
		OutOrderNo:        request.OutOrderNo,
		OutRefundNo:       request.OutRefundNo,
		CpExtra:           request.CpExtra,
		OrderEntrySchema:  request.OrderEntrySchema,
		NotifyUrl:         request.NotifyUrl,
		ItemOrderDetail:   request.ItemOrderDetail,
		RefundTotalAmount: request.RefundTotalAmount,
	}
	marshal, _ := json.Marshal(body)
	path := "/api/apps/trade/v2/create_refund"
	res, err := r.sign(string(marshal), info.ApplicationPrivateKey, path)
	if err != nil {
		return nil, err
	}
	authorization := r.getAuthorization(request.Appid, "1", res)
	out := &douyin.CreateREfundResponse{}
	_ = r.TradeRequest(path, marshal, info.IsSandbox, out, authorization)
	return out, nil
}

func (r *DouyinTrade) QueryRefund(ctx context.Context, request *douyin.QueryRefundRequest) (*douyin.QueryRefundResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}
	body := &douyin.QueryRefundRequest{
		OrderId:     request.OrderId,
		RefundId:    request.RefundId,
		OutRefundNo: request.OutRefundNo,
	}
	marshal, _ := json.Marshal(body)
	path := "/api/apps/trade/v2/query_refund"
	res, err := r.sign(string(marshal), info.ApplicationPrivateKey, path)
	if err != nil {
		return nil, err
	}
	authorization := r.getAuthorization(request.Appid, "1", res)
	out := &douyin.QueryRefundResponse{}
	_ = r.TradeRequest(path, marshal, info.IsSandbox, out, authorization)
	return out, nil
}

func (r *DouyinTrade) QuerySettle(ctx context.Context, request *douyin.QuerySettleRequest) (*douyin.QuerySettleResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}
	body := &douyin.QuerySettleRequest{
		OrderId:     request.OrderId,
		OutOrderNo:  request.OutOrderNo,
		OutSettleNo: request.OutSettleNo,
		SettleId:    request.SettleId,
	}
	marshal, _ := json.Marshal(body)
	path := "/api/apps/trade/v2/query_settle"
	res, err := r.sign(string(marshal), info.ApplicationPrivateKey, path)
	if err != nil {
		return nil, err
	}
	authorization := r.getAuthorization(request.Appid, "1", res)
	out := &douyin.QuerySettleResponse{}
	_ = r.TradeRequest(path, marshal, info.IsSandbox, out, authorization)
	return out, nil
}

func (r *DouyinTrade) QueryOrder(ctx context.Context, request *douyin.QueryOrderRequest) (*douyin.QueryOrderResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}
	body := &douyin.QueryOrderRequest{
		OrderId:    request.OrderId,
		OutOrderNo: request.OutOrderNo,
	}
	marshal, _ := json.Marshal(body)
	path := "/api/apps/trade/v2/query_order"
	res, err := r.sign(string(marshal), info.ApplicationPrivateKey, path)
	if err != nil {
		return nil, err
	}
	authorization := r.getAuthorization(request.Appid, "1", res)
	out := &douyin.QueryOrderResponse{}
	_ = r.TradeRequest(path, marshal, info.IsSandbox, out, authorization)
	return out, nil
}

func (r *DouyinTrade) QueryCps(ctx context.Context, request *douyin.QueryCpsRequest) (*douyin.QueryCpsResponse, error) {
	info, err := config.CFG.GetDouyinAppidInfo(request.Appid)
	if err != nil {
		return nil, err
	}
	body := &douyin.QueryCpsRequest{
		OrderId:    request.OrderId,
		OutOrderNo: request.OutOrderNo,
	}
	marshal, _ := json.Marshal(body)
	path := "/api/apps/trade/v2/query_cps"
	res, err := r.sign(string(marshal), info.ApplicationPrivateKey, path)
	if err != nil {
		return nil, err
	}
	authorization := r.getAuthorization(request.Appid, "1", res)
	out := &douyin.QueryCpsResponse{}
	_ = r.TradeRequest(path, marshal, info.IsSandbox, out, authorization)
	return out, nil
}

func (r *DouyinTrade) TradeRequest(path string, data []byte, IsSandbox bool, out interface{}, authorization string) error {
	host := DOUYIN_API_HOST
	if IsSandbox {
		host = DOUYIN_SANDBOX_API_HOST
	}
	url := host + path
	marshal := data

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Byte-Authorization"] = authorization
	//沙盒环境自动审核参数，REJECT代表审核拒绝，其他代表审核成功
	if IsSandbox {
		headers["Aweme-Check-Type"] = "OK"
	}
	request, err := tools.HttpRequest("POST", url, marshal, headers)
	fmt.Println("request = ", string(request))
	err = json.Unmarshal(request, out)
	if err != nil {
		return err
	}

	return nil
}

// 生成sign
func (r *DouyinTrade) sign(body string, privateKeyStr string, path string) (res map[string]string, e error) {
	res = make(map[string]string)
	// 将私钥字符串解析为PEM块
	pemBlock, _ := pem.Decode([]byte(privateKeyStr))
	if pemBlock == nil {
		return res, errors.New("failed to parse PEM block containing private key")
	}
	// 解析DER编码的私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return res, err
	}
	//生成签名
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce, _ := tools.GenerateRandomString(16)
	sign, e := tools.GenSign("POST", path, timestamp, nonce, body, privateKey)
	if e != nil {
		return res, err
	}
	//签名
	res["sign"] = sign
	//时间戳
	res["timestamp"] = timestamp
	//随机串
	res["nonce"] = nonce
	return
}

// 以上五项签名信息，无顺序要求。按照以下示例格式，key="value"，签名信息之间用英文逗号,隔开。
func (r *DouyinTrade) getAuthorization(appid string, version string, data map[string]string) string {
	return "SHA256-RSA2048 appid=\"" + appid + "\",nonce_str=\"" + data["nonce"] + "\",timestamp=\"" + data["timestamp"] + "\",key_version=\"" + version + "\",signature=\"" + data["sign"] + "\""
}
