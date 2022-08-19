package main

import (
	hello_grpc "base/hello_grpc/pb/hello"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	con, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer con.Close()
	client := hello_grpc.NewHelloGrpcClient(con)
	// 第一种标准
	req, _ := client.Search(context.Background(), &hello_grpc.Req{
		Message: "好久不见",
		Age:     12,
	})
	fmt.Println(req.GetMessage())
	fmt.Println(req.GetAge())

	// 第二钟流式输入
	// Send(*Req) error
	// CloseAndRecv() (*Res, error)
	// grpc.ClientStream
	// c, _ := client.SearchIn(context.Background())
	// i := 0
	// for {
	// 	if i > 10 {
	// 		res, _ := c.CloseAndRecv()
	// 		fmt.Println(res)
	// 		break
	// 	}
	// 	c.Send(&hello_grpc.Req{Message: "我是进来的信息"})
	// 	i++
	// }

	// 第三种流式输出
	// Recv() (*Res, error)
	// grpc.ClientStream
	// c, _ := client.SearchOut(context.Background(), &hello_grpc.Req{
	// 	Message: "好久不见啊，baby!",
	// 	Age:     20,
	// })
	// i := 0
	// for {
	// 	if i > 10 {
	// 		break
	// 	}
	// 	req, _ := c.Recv()
	// 	fmt.Println(req)
	// }

	// 第四种IO输出
	// Send(*Req) error
	// Recv() (*Res, error)
	// grpc.ClientStream
	// c, _ := client.SearchIO(context.Background())
	// wg := &sync.WaitGroup{}
	// wg.Add(2)
	// // 发送消息
	// go func() {
	// 	for {
	// 		err := c.Send(&hello_grpc.Req{
	// 			Message: "hello,world!",
	// 			Age:     20,
	// 		})
	// 		if err != nil {
	// 			wg.Done()
	// 			break
	// 		}
	// 	}
	// }()

	// // 接收消息
	// go func() {
	// 	for {
	// 		res, err := c.Recv()
	// 		if err != nil {
	// 			wg.Done()
	// 			fmt.Println(res)
	// 		}
	// 		fmt.Println(res)
	// 	}
	// }()
	// wg.Wait()

}
