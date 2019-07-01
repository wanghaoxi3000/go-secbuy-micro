package basic

import (
	"github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/basic/config"
	"github.com/wanghaoxi3000/go-secbuy-mirco/stock-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}

func InitSpeciyConfig(path string) {
	config.InitSpecifyFile(path)
	db.Init()
}
