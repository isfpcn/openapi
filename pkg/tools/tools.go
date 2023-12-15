package tools

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

// 签名
func GenSign(method, url, timestamp, nonce, body string, privateKey *rsa.PrivateKey) (string, error) {
	//method内容必须大写，如GET、POST，uri不包含域名，必须以'/'开头
	targetStr := method + "\n" + url + "\n" + timestamp + "\n" + nonce + "\n" + body + "\n"
	h := sha256.New()
	h.Write([]byte(targetStr))
	digestBytes := h.Sum(nil)

	signBytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digestBytes)
	if err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(signBytes)

	return sign, nil
}

// 验证签名
func CheckSign(timestamp, nonce, body, signature, pubKeyStr string) (bool, error) {

	pubKey, err := PemToRSAPublicKey(pubKeyStr) // 注意验签时publicKey使用平台公钥而非应用公钥
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256([]byte(timestamp + "\n" + nonce + "\n" + body + "\n"))
	signBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signBytes)
	return err == nil, nil
}

func PemToRSAPublicKey(pemKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemKeyStr))
	if block == nil || len(block.Bytes) == 0 {
		return nil, fmt.Errorf("empty block in pem string")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch key := key.(type) {
	case *rsa.PublicKey:
		return key, nil
	default:
		return nil, fmt.Errorf("not rsa public key")
	}
}

// HttpRequest 发送请求
func HttpRequest(method string, url string, body []byte, headers map[string]string) ([]byte, error) {

	//fmt.Println("request url =", url)
	//fmt.Println("request body =", string(body))

	buffer := bytes.NewBuffer(body)

	request, err := http.NewRequest(method, url, buffer)

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, val := range headers {
			request.Header.Set(key, val)
		}
	}

	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return all, nil
}

// UrlToFile 图片url转成*os.File
func UrlToFile(url string, ext ...string) (*os.File, error) {
	// 发送 GET 请求获取图片
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fileExt := "jpg"

	if len(ext) > 0 {
		fileExt = ext[0]
	}
	// 创建临时文件
	tmpfile, err := os.CreateTemp("", MD5(url)+"."+fileExt)
	if err != nil {
		return nil, err
	}

	// 将图片内容写入到临时文件中
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		return nil, err
	}

	// 将文件指针重置到文件开始处
	_, err = tmpfile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return tmpfile, nil
}

func MD5(str string) string {
	harsher := md5.New()
	harsher.Write([]byte(str))
	hashInBytes := harsher.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString
}

func UploadFile(file *os.File, url string, fileFieldName string, extraFields map[string]string) ([]byte, error) {
	// Create a buffer to store the request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	defer file.Close()
	fileField, err := writer.CreateFormFile(fileFieldName, file.Name())
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileField, file)
	if err != nil {
		return nil, err
	}

	if len(extraFields) > 0 {
		for k, v := range extraFields {
			_ = writer.WriteField(k, v)
		}
	}

	// Close the multipart writer
	_ = writer.Close()

	// Create the HTTP request
	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return nil, err
	}

	// Set the content type
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	all, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return all, err
}
