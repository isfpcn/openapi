/**
 * @author   Liang
 * @create   2023/11/30 18:35
 * @version  1.0
 */

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"openapi/api/douyin"
)

func main() {

	// 创建 gRPC 连接
	conn, err := grpc.Dial(":8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := douyin.NewAccessTokenClient(conn)

	req := &douyin.GetAccessTokenRequest{
		Appid:     "tt39164faba6284b0001y",
		GrantType: "client_credential",
	}

	material, err := client.GetAccessToken(context.Background(), req)

	fmt.Println("err = ", err)
	fmt.Println("material = ", material)

	st, ok := status.FromError(err)
	if ok {
		fmt.Printf("gRPC status code: %v\n", st.Code().String())
		fmt.Printf("Error message: %v\n", st.Message())
	} else {
		fmt.Printf("Non-gRPC error: %v\n", err)
	}

}
