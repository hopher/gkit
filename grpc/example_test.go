package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func Example_basic() {

	errChan := make(chan error)

	// 初始化 GRPC 服务
	gs := grpc.NewServer(
		grpc.UnaryInterceptor(ChainUnaryServer(
			ErrorHandleInterceptor(nil),
		)),
		grpc.StreamInterceptor(ChainStreamServer(
		// ...
		)),
	)

	// GRPC 服务
	go (func() {
		ls, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
		if err != nil {
			errChan <- err
		}
		errChan <- gs.Serve(ls)
	})()

	// 退出信号处理
	go (func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-ch)
	})()

	// 通道阻塞
	<-errChan
}
