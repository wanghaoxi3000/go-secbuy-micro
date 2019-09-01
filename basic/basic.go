package basic

import (
	"github.com/wanghaoxi3000/go-secbuy-micro/basic/config"
	"github.com/wanghaoxi3000/go-secbuy-micro/basic/db"
)

// Init 初始化配置，读取默认的配置路径
func Init() {
	config.Init()
	db.Init()
}

// InitSpeciyConfig 初始化配置，读取指定的配置路径
func InitSpeciyConfig(path string) {
	config.InitSpecifyFile(path)
	db.Init()
}
