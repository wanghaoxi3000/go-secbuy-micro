package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	stock "github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/proto/stock"
)

var (
	serviceClient stock.StockService
)

// Error 错误结构体
type Error struct {
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = stock.NewStockService("go.micro.secbuy.srv.stock", client.DefaultClient)
}

// CommodityInfo 查询 Commodity 信息
func CommodityInfo(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	id := int32(request["id"].(float64))
	log.Logf("receive ID %d", id)
	rsp, err := serviceClient.GetCommodity(context.TODO(), &stock.GetRequest{
		Id: id,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	log.Logf("query result: success %v, error: [%v] %v",
		rsp.GetSuccess(), rsp.GetError(), rsp.GetError())
	response := map[string]interface{}{
		"success":   rsp.GetSuccess(),
		"commodity": rsp.GetCommodity(),
	}
	if rsp.GetSuccess() {
		response["error"] = nil
	} else {
		response["error"] = &Error{
			Code:   rsp.GetError().Code,
			Detail: rsp.GetError().Detail,
		}
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
