package basic

import (
	config2 "github.com/wanghaoxi3000/go-secbuy-mirco/basic/config"
	"github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/basic/db"
)

func Init() {
	config2.Init()
	db.Init()
}

func InitSpeciyConfig(path string) {
	config2.InitSpecifyFile(path)
	db.Init()
}
