package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/wanghaoxi3000/go-secbuy-mirco/basic"
	"github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv/handler"
	"github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv/model"
	payment "github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv/proto/payment"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.secbuy.srv.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// Register Handler
	paymentHandle := new(handler.Payment)
	micro.RegisterSubscriber("payment.payevent", service.Server(), paymentHandle.Process)
	payment.RegisterPaymentServiceHandler(service.Server(), paymentHandle)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
