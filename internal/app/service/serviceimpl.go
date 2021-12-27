package service

import (
	"errors"
	"fibonacciServiceGRPC/proto"
)

type ServiceImpl struct {}

func NewService() Service {
	return &ServiceImpl{}
}

func (s ServiceImpl) GetFibonacciList(req *proto.GetFibonacciSliceReq) (*proto.GetFibonacciSliceRes, error) {
	if req.Y > 95{
		return nil, errors.New("too big value for y")
	} else if req.X > req.Y{
		return nil, errors.New("x should be more than y")
	}

	f := fib(req.X)
	res := &proto.GetFibonacciSliceRes{}
	var i uint64
	for i = 0; i < req.Y; i++{
		res.Res = append(res.Res, f())
	}

	return res, nil
}

func fib(from uint64) func() uint64{
	x := from
	y := from+1

	return func() uint64 {
		x, y = y, x + y
		return x
	}
}