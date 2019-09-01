package stock

import (
	"testing"

	"github.com/wanghaoxi3000/go-secbuy-micro/basic"
	proto "github.com/wanghaoxi3000/go-secbuy-micro/stock-srv/proto/stock"
)

func TestCreateCommodity(t *testing.T) {
	basic.InitSpeciyConfig("../../conf")
	Init()

	c := &proto.Commodity{
		Name:  "测试",
		Count: 10,
		Sale:  0,
	}
	s.CreateCommodity(c)
}
