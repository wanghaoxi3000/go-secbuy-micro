package db

import (
	"fmt"
	"sync"

	"github.com/wanghaoxi3000/go-secbuy-mirco/basic/config"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
)

var (
	inited bool
	db     *gorm.DB
	m      sync.RWMutex
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Logf(err.Error())
		return
	}

	// 如果配置声明使用mysql
	if config.GetPostgresConfig().GetEnabled() {
		db = initPostgres()
	}

	inited = true
}

// GetDB 获取db
func GetDB() *gorm.DB {
	return db
}
