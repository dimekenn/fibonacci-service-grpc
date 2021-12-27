package service

import "fibonacciServiceGRPC/proto"

type Service interface {
	GetFibonacciList(req *proto.GetFibonacciSliceReq) (*proto.GetFibonacciSliceRes, error)
}