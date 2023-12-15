package http

import (
	"errors"
	"google.golang.org/grpc/status"
	"openapi/api/douyin"
	code2 "openapi/internal/code"
	"openapi/internal/config"
	"openapi/internal/repo"
	douyin2 "openapi/internal/service/douyin"
	"openapi/internal/transport/http/md"
	"openapi/internal/validate"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gogo/protobuf/proto"
)

type (
	httpError struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	httpSuccess struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data"`
	}

	grpcStatus interface {
		GRPCStatus() *status.Status
	}
)

var (

	// WriterHandler define  writer handler
	WriterHandler = func(c *fiber.Ctx, resp proto.Message) error {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		return c.JSON(httpSuccess{Msg: "ok", Data: resp})
	}

	// DefaultErrorHandler Default error handler
	DefaultErrorHandler = func(c *fiber.Ctx, err error) error {

		// Status code defaults to 100
		code := code2.InvalidArgument
		msg := err.Error()

		var e *fiber.Error
		var gs grpcStatus

		if errors.As(err, &e) {
			code = code2.InvalidArgument
		} else if errors.As(err, &gs) {
			st, ok := status.FromError(err)
			if ok {
				msg = st.Message()
				code = st.Code()
			}
		}

		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		// Return status code with error message
		return c.JSON(httpError{Code: int(code), Msg: msg})
	}
)

func New() error {

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: DefaultErrorHandler,
		//打印路由
		EnablePrintRoutes: config.CFG.Server.HttpDumpRoute,

		//关闭fiber服务启动信息输出
		DisableStartupMessage: true,

		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(md.Log)

	initRouter(app)

	douyinRepo := new(repo.DouyinRepo)
	douyin.RegisterDouyinTradeFiberServer(app, &douyin2.TradeService{Repo: douyinRepo}, WriterHandler, validate.ValidateHandler)
	douyin.RegisterDouyinClientTokenFiberServer(app, &douyin2.ClientTokenService{Repo: douyinRepo}, WriterHandler, validate.ValidateHandler)
	douyin.RegisterDouyinAuthFiberServer(app, &douyin2.AuthService{Repo: douyinRepo}, WriterHandler, validate.ValidateHandler)
	douyin.RegisterDouyinAccessTokenFiberServer(app, &douyin2.AccessTokenService{Repo: douyinRepo}, WriterHandler, validate.ValidateHandler)
	douyin.RegisterDouyinProductFiberServer(app, &douyin2.ProductService{Repo: douyinRepo}, WriterHandler, validate.ValidateHandler)
	err := app.Listen(config.CFG.Server.HttpAddr)
	return err
}

func initRouter(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		_, err := ctx.WriteString("Version: " + config.CFG.Version)
		return err
	})
}
