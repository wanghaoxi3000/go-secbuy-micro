package basic

import (
	"github.com/wanghaoxi3000/go-secbuy-mirco/basic/config"
	"github.com/wanghaoxi3000/go-secbuy-mirco/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
