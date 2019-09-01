module github.com/wanghaoxi3000/go-secbuy-micro/payment-srv

go 1.12

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.8.3
	github.com/nats-io/nats-server/v2 v2.0.4 // indirect
	github.com/wanghaoxi3000/go-secbuy-micro/basic v0.0.0-00010101000000-000000000000
)

replace github.com/wanghaoxi3000/go-secbuy-micro/basic => ../basic
