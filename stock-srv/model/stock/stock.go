package stock

import (
	"fmt"
	"sync"
	"time"

	"github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/basic/db"
	proto "github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/proto/stock"
)

var (
	s *service
	m sync.RWMutex
)

// Service 仓库服务类
type Service interface {
	CreateCommodity(commodity *proto.Commodity) (err error)
	QueryCommodityByID(id int32) (ret *proto.Commodity, err error) // QueryCommodityByID 根据ID获取商品信息
}

// service 服务
type service struct {
}

type stockModel struct {
	ID         int32
	Name       string
	Count      int32
	Sale       int32
	CreateTime time.Time `gorm:"DEFAULT:now()"`
}

func (stockModel) TableName() string {
	return "stock"
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

func (s *service) CreateCommodity(commodity *proto.Commodity) error {
	o := db.GetDB()

	model := stockModel{
		Name:  commodity.GetName(),
		Count: commodity.GetCount(),
		Sale:  commodity.GetSale(),
	}
	o.Create(&model)
	return nil
}

func (s *service) QueryCommodityByID(id int32) (*proto.Commodity, error) {
	o := db.GetDB()

	model := &stockModel{}
	o.Where("id = ?", id).First(model)
	fmt.Println(model.CreateTime.Local(), model.CreateTime.UTC())
	commodity := &proto.Commodity{
		Id:         model.ID,
		Name:       model.Name,
		Count:      model.Count,
		Sale:       model.Sale,
		CreateTime: model.CreateTime.Format("2006-01-02T15:04:05Z07:00"),
	}

	return commodity, nil
}
