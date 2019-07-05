package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	order "github.com/wanghaoxi3000/go-secbuy-mirco/order-srv/proto/order"
)

var (
	serviceClient order.OrderService
)

// Error 错误结构体
type Error struct {
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = order.NewOrderService("go.micro.secbuy.srv.order", client.DefaultClient)
}

// Sell 订购接口
func Sell(w http.ResponseWriter, r *http.Request) {
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
	log.Logf("receive sell ID %d", id)
	rsp, err := serviceClient.CreateOrder(context.TODO(), &order.GetRequest{
		Id: id,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	log.Logf("sell result: success %v, error: [%v] %v",
		rsp.GetSuccess(), rsp.GetError(), rsp.GetError())
	response := map[string]interface{}{
		"success": rsp.GetSuccess(),
	}
	if rsp.GetSuccess() {
		response["order"] = rsp.GetOrder()
		response["error"] = nil
	} else {
		response["order"] = nil
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
