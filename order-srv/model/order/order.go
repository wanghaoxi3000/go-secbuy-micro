package stock

import (
	"fmt"
	"sync"
	"time"

	proto "github.com/wanghaoxi3000/go-secbuy-mirco/order-srv/proto/order"
	"github.com/wanghaoxi3000/go-secbuy-mirco/basic/db"
)

var (
	s *service
	m sync.RWMutex
)

// OrderService 订单服务
type OrderService interface {
	CreateOrder(int32) proto.Order error
}

type service struct {
}

type orderModel struct {
	ID         int32
	Sid		   int32
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

	s = &service{}
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

func (s *service) CreateOrder() (proto.Order error){
	
}

func (s *service) CreateCommodity(commodity *proto.Commodity) error {
	o := db.GetDB()

	model := stockModel{
		Name:  commodity.GetName(),
		Count: commodity.GetCount(),
		Sale:  commodity.GetSale(),
	}
	o.Create(&model)
	commodity.Id = model.ID
	commodity.CreateTime = model.CreateTime.Format("2006-01-02T15:04:05")
	return nil
}

func (s *service) QueryCommodityByID(id int32) (*proto.Commodity, error) {
	o := db.GetDB()

	model := &stockModel{}
	if err := o.Where("id = ?", id).First(model).Error; err != nil {
		return nil, err
	}

	commodity := &proto.Commodity{
		Id:         model.ID,
		Name:       model.Name,
		Count:      model.Count,
		Sale:       model.Sale,
		CreateTime: model.CreateTime.Format("2006-01-02T15:04:05"),
	}

	return commodity, nil
}
