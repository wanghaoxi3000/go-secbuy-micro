module github.com/wanghaoxi3000/go-secbuy-mirco/order-web

go 1.12

require (
	github.com/micro/go-micro v1.8.3
	github.com/wanghaoxi3000/go-secbuy-mirco/order-srv v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/wanghaoxi3000/go-secbuy-mirco/basic => ../basic

replace github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv => ../stock-srv

replace github.com/wanghaoxi3000/go-secbuy-mirco/order-srv => ../order-srv
