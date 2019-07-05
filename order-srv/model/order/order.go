package order

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	"github.com/wanghaoxi3000/go-secbuy-mirco/basic/db"
	proto "github.com/wanghaoxi3000/go-secbuy-mirco/order-srv/proto/order"
	stockSrv "github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/proto/stock"
)

var (
	s           *service
	m           sync.RWMutex
	stockClient stockSrv.StockService
)

// Service 订单服务
type Service interface {
	CreateOrder(int32) (*proto.Order, error)
}

type service struct {
}

type order struct {
	ID         int32
	Sid        int32
	Name       string
	CreateTime time.Time `gorm:"DEFAULT:now()"`
}

// Init 初始化库存服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	stockClient = stockSrv.NewStockService("go.micro.secbuy.srv.stock", client.DefaultClient)
	s = &service{}
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

func (s *service) CreateOrder(id int32) (*proto.Order, error) {
	rsp, err := stockClient.Sell(context.TODO(), &stockSrv.GetRequest{Id: id})
	if err != nil {
		log.Logf("[model] Sell 调用库存服务时失败：%s", err.Error())
		return nil, err
	}
	if !rsp.GetSuccess() {
		return nil, errors.New("销存失败")
	}

	o := db.GetDB()
	model := order{
		Sid:  id,
		Name: rsp.GetCommodity().GetName(),
	}
	o.Create(&model)
	orderProto := &proto.Order{
		Id:         model.ID,
		Name:       model.Name,
		CreateTime: model.CreateTime.Format("2006-01-02T15:04:05"),
	}

	return orderProto, nil
}
