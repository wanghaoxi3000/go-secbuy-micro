package handler

import (
	"context"
	"github.com/micro/go-micro"

	"github.com/micro/go-micro/util/log"

	order "github.com/wanghaoxi3000/go-secbuy-mirco/order-srv/model/order"
	proto "github.com/wanghaoxi3000/go-secbuy-mirco/order-srv/proto/order"
	paymentProto "github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv/proto/payment"
)

var (
	orderModel order.Service
)

type Order struct{
	PaymentPublisher    micro.Publisher
}

// Init 初始化handler
func Init() {
	var err error
	orderModel, err = order.GetService()
	if err != nil {
		log.Fatalf("[Init] 初始化Handler错误: %s", err.Error())
		return
	}
}

// CreateOrder 创建订单
func (e *Order) CreateOrder(ctx context.Context, req *proto.GetRequest, rsp *proto.Response) error {
	log.Log("Received Order.CreateOrder request")
	protoOrder, err := orderModel.CreateOrder(req.Id)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{Code: 400, Detail: err.Error()}
		return nil
	}
	log.Logf("[handle] publish create payment message")
	e.PaymentPublisher.Publish(ctx, &paymentProto.PayEvent{Id: protoOrder.Id, Name: protoOrder.Name})
	rsp.Success = true
	rsp.Order = protoOrder
	return nil
}
