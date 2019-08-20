module github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10 // indirect
	github.com/micro/go-micro v1.8.3
	github.com/micro/protoc-gen-micro v0.8.0 // indirect
	github.com/wanghaoxi3000/go-secbuy-mirco/basic v0.0.0-00010101000000-000000000000
)

replace github.com/wanghaoxi3000/go-secbuy-mirco/basic => ../basic
