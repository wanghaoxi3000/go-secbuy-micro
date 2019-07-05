module github.com/wanghaoxi3000/go-secbuy-mirco/order-srv

go 1.12

require (
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-micro v1.7.0
	github.com/wanghaoxi3000/go-secbuy-mirco/basic v0.0.0-00010101000000-000000000000
	github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv v0.0.0-00010101000000-000000000000
)

replace github.com/wanghaoxi3000/go-secbuy-mirco/basic => ../basic

replace github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv => ../stock-srv
