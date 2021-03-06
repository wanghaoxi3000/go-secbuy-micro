package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	modelStock "github.com/wanghaoxi3000/go-secbuy-micro/stock-srv/model/stock"
	stock "github.com/wanghaoxi3000/go-secbuy-micro/stock-srv/proto/stock"
)

var (
	stockModel modelStock.Service
)

type Stock struct{}

// Init 初始化handler
func Init() {
	var err error
	stockModel, err = modelStock.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// CreateCommodity 创建商品
func (e *Stock) CreateCommodity(ctx context.Context, req *stock.Commodity, rsp *stock.Response) error {
	log.Log("Received Stock.CreateCommodity request")
	stockModel.CreateCommodity(req)
	rsp.Success = true
	rsp.Commodity = req
	return nil
}

// GetCommodity 根据 ID 查询商品
func (e *Stock) GetCommodity(ctx context.Context, req *stock.GetRequest, rsp *stock.Response) error {
	log.Logf("Received Stock.GetCommodity request with ID: %d", req.Id)

	commodity, err := stockModel.QueryCommodityByID(req.Id)
	if err != nil {
		rsp.Success = false
		rsp.Error = &stock.Error{
			Code:   404,
			Detail: err.Error(),
		}
		return nil
	}

	rsp.Success = true
	rsp.Commodity = commodity
	return nil
}

// Sell 根据 ID 销存
func (e *Stock) Sell(ctx context.Context, req *stock.GetRequest, rsp *stock.Response) error {
	log.Logf("Received Stock.Sell request with ID: %d", req.Id)
	commodity, err := stockModel.SellCommodityByID(req.Id)
	if err != nil {
		rsp.Success = false
		rsp.Error = &stock.Error{
			Code:   400,
			Detail: err.Error(),
		}
		return nil
	}

	rsp.Success = true
	rsp.Commodity = commodity
	return nil
}
