package stock

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/micro/go-micro/util/log"

	"github.com/wanghaoxi3000/go-secbuy-mirco/basic/db"
	proto "github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/proto/stock"
)

var (
	s *service
	m sync.RWMutex
)

// Service 仓库服务类
type Service interface {
	CreateCommodity(commodity *proto.Commodity) error
	QueryCommodityByID(id int32) (*proto.Commodity, error) // QueryCommodityByID 根据ID获取商品信息
	SellCommodityByID(id int32) (*proto.Commodity, error)
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
	Version    int32
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

// CreateCommodity 创建商品库存
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

// QueryCommodityByID 查询商品库存
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

// Sell 销存
func (s *service) SellCommodityByID(id int32) (commodity *proto.Commodity, err error) {
	o := db.GetDB()
	model := &stockModel{}
	if err := o.Where("id = ?", id).First(model).Error; err != nil {
		log.Logf("Find %d commodity error: %v", id, err)
		return nil, err
	}
	log.Logf("Sell id: %d name: %s version: %d", model.ID, model.Name, model.Version)

	model.Sale++
	if model.Sale > model.Count {
		err = errors.New("commodity sales complete")
		return
	}

	if row := o.Model(&model).Where("version = ?", model.Version).Updates(
		map[string]interface{}{
			"sale":    model.Sale,
			"version": model.Version + 1,
		}).RowsAffected; row == 0 {
		return nil, errors.New("commodity info timeout")
	}

	commodity = &proto.Commodity{
		Id:         model.ID,
		Name:       model.Name,
		Count:      model.Count,
		Sale:       model.Sale,
		CreateTime: model.CreateTime.Format("2006-01-02T15:04:05"),
	}
	return
}
