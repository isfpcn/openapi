package config

import "errors"

var CFG *Config

type Server struct {
	HttpAddr      string `json:"http_addr"`
	HttpDumpRoute bool   `json:"http_dump_route"`
	GrpcAddr      string `json:"grpc_addr"`
}

type Log struct {
	LogFile string `json:"log_file"`
}

type DataBase struct {
	DSN             string `json:"dsn"`
	MaxIdleConn     int    `json:"max_idle_conn"`
	MaxOpenConn     int    `json:"max_open_conn"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"`
	Debug           bool   `json:"debug"`
}

type Redis struct {
	Address  []string `json:"address"`
	Password string   `json:"password"`
}

type Douyin struct {
	Appids []DouyinAppidInfo `json:"appids"`
}

type DouyinAppidInfo struct {
	Appid                 string `json:"appid"`
	Secret                string `json:"secret"`
	IsSandbox             bool   `json:"is_Sandbox"`
	PlatformPublicKey     string `json:"platform_public_key"`
	ApplicationPrivateKey string `json:"application_private_key"`
}

type Config struct {
	Version  string
	Env      string
	Server   Server   `json:"server"`
	Log      Log      `json:"log"`
	DataBase DataBase `json:"dataBase"`
	Redis    Redis    `json:"redis"`
	Douyin   Douyin   `json:"douyin"`
}

func (c *Config) GetDouyinAppidInfo(appid string) (*DouyinAppidInfo, error) {

	for _, v := range c.Douyin.Appids {
		if v.Appid == appid {
			return &v, nil
		}
	}
	return nil, errors.New("抖音appid信息未配置")
}
