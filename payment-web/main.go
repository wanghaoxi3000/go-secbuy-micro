package main

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/go-micro/util/log"

	"github.com/wanghaoxi3000/go-secbuy-mirco/payment-web/handler"

	"github.com/micro/go-micro/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.secbuy.web.payment"),
		web.Version("latest"),
		web.Address(":8088"),
	)

	// initialise service
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/payment/payorder", handler.PayOrder)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
