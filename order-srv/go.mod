module github.com/wanghaoxi3000/go-secbuy-micro/order-srv

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.8.3
	github.com/wanghaoxi3000/go-secbuy-micro/basic v0.0.0-00010101000000-000000000000
	github.com/wanghaoxi3000/go-secbuy-micro/payment-srv v0.0.0-00010101000000-000000000000
	github.com/wanghaoxi3000/go-secbuy-micro/stock-srv v0.0.0-00010101000000-000000000000
)

replace github.com/wanghaoxi3000/go-secbuy-micro/basic => ../basic

replace github.com/wanghaoxi3000/go-secbuy-micro/stock-srv => ../stock-srv

replace github.com/wanghaoxi3000/go-secbuy-micro/payment-srv => ../payment-srv
