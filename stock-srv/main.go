package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/wanghaoxi3000/go-secbuy-mirco/basic"
	"github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/handler"
	"github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/model"
	stock "github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/proto/stock"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.secbuy.srv.stock"),
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
	stock.RegisterStockServiceHandler(service.Server(), new(handler.Stock))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
