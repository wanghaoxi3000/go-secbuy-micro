module github.com/wanghaoxi3000/go-secbuy-mirco/payment-web

go 1.12

replace github.com/wanghaoxi3000/go-secbuy-mirco/basic => ../basic

replace github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv => ../payment-srv

require (
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.9.1
	github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv v0.0.0-00010101000000-000000000000
)
