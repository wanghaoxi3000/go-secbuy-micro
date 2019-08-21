package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	payment "github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv/model/payment"
	proto "github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv/proto/payment"
)

var (
	paymentService payment.Service
)

type Payment struct{}

// Init 初始化handler
func Init() {
	var err error
	paymentService, err = payment.GetService()
	if err != nil {
		log.Fatalf("[Init] 初始化Handler错误: %s", err.Error())
		return
	}
}

// Process 接收创建支付信息消息
func (e *Payment) Process(ctx context.Context, event *proto.PayEvent) error {
	log.Log("Received create payment message")
	if err := paymentService.CreatePayment(event); err!= nil{
		return err
	}

	return nil
}

// PayOrder 支付订单
func (e *Payment) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Log("Received Payment.PayOrder request")
	payment.GetService()
	paymentInfo, err := paymentService.PayOrder(req.Id)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{Code: 400, Detail: err.Error()}
		return nil
	}
	rsp.Success = true
	rsp.Payment = paymentInfo
	return nil
}
