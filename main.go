package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"openapi/internal/config"
	"openapi/internal/dao"
	"openapi/internal/logger"
	"openapi/internal/transport/grpc"
	"openapi/internal/transport/http"
	"openapi/internal/validate"
	"os"
	"os/signal"
	"syscall"

	oslog "log"

	"github.com/spf13/viper"
)

var (
	genModel bool
	Version  string
)

func init() {
	flag.BoolVar(&genModel, "gen_model", false, "gen gorm model")
}

func main() {

	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			if logger.Logger != nil {
				logger.Logger.Error("%v", err)
			}
		}
		dao.Close()
	}()

	v := viper.New()

	v.SetConfigType("toml")
	v.SetConfigFile("config/config.toml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	cb, err := json.Marshal(v.AllSettings())

	if err != nil {
		panic(fmt.Errorf("Fatal error config json.Marshal: %s \n", err))
	}

	conf := &config.Config{}
	err = json.Unmarshal(cb, conf)

	if err != nil {
		panic(fmt.Errorf("Fatal error config json.Unmarshal: %s \n", err))
	}

	config.CFG = conf
	config.CFG.Version = Version

	logger.InitLogger()

	db, cf, err := dao.NewDB()

	if err != nil {
		panic(fmt.Errorf("Fatal error NewDB: %s \n", err))
	}

	redis, cf2, err := dao.NewRedis()

	if err != nil {
		cf()
		panic(fmt.Errorf("Fatal error NewRedis: %s \n", err))
	}

	dao.NewDao(db, redis)

	if genModel {
		dao.GenModel()
		os.Exit(0)
	}

	err = validate.New()

	if err != nil {
		panic(fmt.Errorf("Fatal error validate new: %s \n", err))
	}

	go func() {
		oslog.Printf("http server start port is %s", config.CFG.Server.HttpAddr)
		err = http.New()
		if err != nil {
			cf()
			cf2()
			panic(fmt.Errorf("Fatal error http start: %s \n", err))
		}
	}()

	go func() {
		oslog.Printf("grpc server start port is %s", config.CFG.Server.GrpcAddr)
		err = grpc.New()
		if err != nil {
			cf()
			cf2()
			panic(fmt.Errorf("Fatal error grpc start: %s \n", err))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		oslog.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			dao.Close()
			oslog.Printf("go exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
