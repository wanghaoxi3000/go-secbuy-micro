package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // import the database’s driver
	"github.com/micro/go-micro/util/log"

	"github.com/wanghaoxi3000/go-secbuy-mirco/basic/config"
)

func initPostgres() *gorm.DB {
	postgresConfig := config.GetPostgresConfig()
	connConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		postgresConfig.GetHost(),
		postgresConfig.GetUser(),
		postgresConfig.GetDBname(),
		postgresConfig.GetPassword(),
	)

	db, err := gorm.Open("postgres", connConfig)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db.DB().Ping()

	return db
}
