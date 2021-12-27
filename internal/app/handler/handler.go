package handler

import (
	"context"
	"fibonacciServiceGRPC/internal/app/service"
	"fibonacciServiceGRPC/proto"
	"github.com/micro/go-micro/util/log"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetFibonacciSlice(ctx context.Context, req *proto.GetFibonacciSliceReq) (*proto.GetFibonacciSliceRes, error) {
	res, err := h.service.GetFibonacciList(req)
	if err != nil{
		log.Errorf("failed to calculate fibonacci: %v", err)
		return nil, err
	}
	return res, nil
}
