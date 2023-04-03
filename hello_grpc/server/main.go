package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	hello_grpc "base/hello_grpc/pb/hello"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type server struct {
	hello_grpc.UnimplementedHelloGrpcServer
}

func (s *server) Search(ctx context.Context, req *hello_grpc.Req) (*hello_grpc.Res, error) {
	message := req.GetMessage()
	age := req.GetAge()
	fmt.Println(req.GetMessage())
	fmt.Println(req.GetAge())
	return &hello_grpc.Res{
		Message: "从客户端接收到的消息为" + message,
		Age:     age + 10,
	}, nil
}

func (s *server) SearchIn(server hello_grpc.HelloGrpc_SearchInServer) error {
	// SendAndClose(*Res) error
	// Recv() (*Req, error) 不间断读取
	// grpc.ServerStream
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&hello_grpc.Res{
				Message: "完成了",
			})
			break
		}
	}
	return nil
}

func (s *server) SearchOut(req *hello_grpc.Req, server hello_grpc.HelloGrpc_SearchOutServer) error {
	fmt.Println(req.GetMessage())
	fmt.Println(req.GetAge())
	message := req.GetMessage()
	age := req.GetAge()
	i := 0
	for {
		if i > 10 {
			break
		}
		server.Send(&hello_grpc.Res{ //nolint:errcheck
			Message: message + "lallalalllll",
			Age:     age + 100,
		})
		i++
	}
	return nil
}

func (s *server) SearchIO(server hello_grpc.HelloGrpc_SearchIOServer) error {
	// Send(*Res) error
	// Recv() (*Req, error)
	// grpc.ServerStream

	str := make(chan string)

	// 流式获取，收到的数据
	go func() {
		for {
			req, err := server.Recv()
			fmt.Println(req)
			if err != nil {
				str <- "结束"
				break
			}
			str <- req.Message
		}
	}()
	for {
		s := <-str
		if s == "结束" {
			break
		}
		// 发出的数据
		server.Send(&hello_grpc.Res{ //nolint:errcheck
			Message: s + "1111111111",
			Age:     1111,
		})
	}

	// 流式获取，收到的数据
	// go func() {
	// 	for {
	// 		req, err := server.Recv()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		message := req.GetMessage()
	// 		age := req.GetAge()
	// 		// 发出的数据
	// 		server.Send(&hello_grpc.Res{
	// 			Message: message + "1111111111",
	// 			Age:     1111 + age,
	// 		})
	// 	}
	// }()

	return nil
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go registerGataway(wg)
	go registerGrpc(wg)
	wg.Wait()
}

func registerGataway(wg *sync.WaitGroup) {
	// 启动后访问 http://localhost:8090/api/search
	// 请求参数 {"message":"xdsa","age":2}
	con, _ := grpc.DialContext(context.Background(), "localhost:8888", grpc.WithBlock(), grpc.WithInsecure())
	defer con.Close()

	mux := runtime.NewServeMux() // 一个对外开放的mux

	gwServer := &http.Server{
		Handler: mux,
		Addr:    ":8090",
	}
	// 注册网关
	err := hello_grpc.RegisterHelloGrpcHandler(context.Background(), mux, con)
	if err != nil {
		fmt.Println(err)
	}
	if err = gwServer.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
	wg.Done()
}

func registerGrpc(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGrpcServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		fmt.Println(err.Error())
	}
	wg.Done()
}
