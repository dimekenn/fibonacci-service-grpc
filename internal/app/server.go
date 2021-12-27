package app

import (
	"fibonacciServiceGRPC/configs"
	"fibonacciServiceGRPC/internal/app/handler"
	"fibonacciServiceGRPC/internal/app/service"
	"fibonacciServiceGRPC/proto"
	"github.com/micro/go-micro/util/log"
	"google.golang.org/grpc"
	"net"
)

func StartGRPCServer(errCh chan<- error, cfg *configs.Configs){
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil{
		errCh <- err
	}
	s:= grpc.NewServer()
	fibService := service.NewService()
	fibHandler := handler.NewHandler(fibService)
	proto.RegisterFibonacciServiceServer(s, fibHandler)

	log.Infof("grpc server started on :%s", cfg.Port)

	errCh <- s.Serve(lis)
}
