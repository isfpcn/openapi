/**
 * @author   Liang
 * @create   2023/11/30 11:55
 * @version  1.0
 */

package md

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"openapi/internal/logger"
	"openapi/internal/validate"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/gogo/protobuf/proto"

	"google.golang.org/grpc"
)

// MiddlewareInterceptor 是一个gRPC中间件的实现
func MiddlewareInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 从上下文中获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// 输出所有元数据
		for key, values := range md {
			//fmt.Printf("Metadata Key: %s, Values: %s\n", key, values)
			ctx = context.WithValue(ctx, key, values[0])
		}
	}

	source := ctx.Value("source")
	track_id := ctx.Value("track-id")

	if track_id == nil {
		seed := uuid.NewV4()
		UUID := uuid.NewV5(seed, "track-id").String()
		track_id = UUID
	}
	startTime := time.Now()

	// 调用处理程序之前的逻辑
	logger.Logger.Infof("RPC Source: %v, TrackId: %v, Method: %s, started at: %s, Request: %v\n", source, track_id, info.FullMethod, startTime.Format(time.RFC3339), req)

	err := validate.ValidateHandler(req.(proto.Message))

	if err != nil {
		return nil, err
	}

	// 调用处理程序
	resp, err := handler(ctx, req)
	// 调用处理程序之后的逻辑
	logger.Logger.Infof("RPC Source: %v, TrackId: %v, Method: %s, completed in: %s Response: %v, Error: %v\n", source, track_id, info.FullMethod, time.Since(startTime), resp, err)

	return resp, err
}
