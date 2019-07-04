package db

import (
	"fmt"
	config2 "github.com/wanghaoxi3000/go-secbuy-mirco/basic/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // import the database’s driver
	"github.com/micro/go-micro/util/log"
)

func initPostgre() *gorm.DB {
	postgreConfig := config2.GetPostgreConfig()
	connConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		postgreConfig.GetHost(),
		postgreConfig.GetUser(),
		postgreConfig.GetDBname(),
		postgreConfig.GetPassword(),
	)

	db, err := gorm.Open("postgres", connConfig)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db.DB().Ping()

	return db
}
