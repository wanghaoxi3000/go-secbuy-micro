package payment

import (
	"fmt"
	"sync"
	"time"

	"github.com/micro/go-micro/util/log"

	"github.com/wanghaoxi3000/go-secbuy-micro/basic/db"
	proto "github.com/wanghaoxi3000/go-secbuy-micro/payment-srv/proto/payment"
)

var (
	s *service
	m sync.RWMutex
)

// Payment 订单服务
type Service interface {
	CreatePayment(event *proto.PayEvent) error
	PayOrder(int32) (*proto.Payment, error)
}

type service struct {
}

type payment struct {
	ID         int32
	Sid        int32
	Name       string
	State      int32
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

func (s *service) CreatePayment(event *proto.PayEvent) error {
	log.Infof("create payment in db, sid: %d name: %s", event.Id, event.Name)
	o := db.GetDB()
	model := &payment{
		Sid:  event.Id,
		Name: event.Name,
	}
	if err := o.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (s *service) PayOrder(id int32) (*proto.Payment, error) {
	log.Infof("pay for order: %d", id)
	o := db.GetDB()

	model := &payment{}
	if err := o.Where("id = ?", id).First(model).Error; err != nil {
		return nil, err
	}

	if err := o.Model(model).Update("state", 1).Error; err != nil {
		return nil, err
	}

	payment := &proto.Payment{
		Id:         model.ID,
		Name:       model.Name,
		State:      1,
		CreateTime: model.CreateTime.Format("2006-01-02T15:04:05"),
	}

	return payment, nil
}
