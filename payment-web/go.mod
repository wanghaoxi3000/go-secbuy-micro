module github.com/wanghaoxi3000/go-secbuy-micro/payment-web

go 1.12

replace github.com/wanghaoxi3000/go-secbuy-micro/basic => ../basic

replace github.com/wanghaoxi3000/go-secbuy-micro/payment-srv => ../payment-srv

require (
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.9.1
	github.com/wanghaoxi3000/go-secbuy-micro/payment-srv v0.0.0-00010101000000-000000000000
)
