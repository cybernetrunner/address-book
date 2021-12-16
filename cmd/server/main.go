package main

import (
	"github.com/cyneruxyz/address-book/internal/app"
	"github.com/cyneruxyz/address-book/internal/database"
	"github.com/cyneruxyz/address-book/pkg/util"
	"github.com/golang/glog"
	"github.com/profclems/go-dotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	errConf = "Fatal error loading .env file: %s \n"
	errDB   = "Fatal error database: %s \n"
)

func main() {
	// util initialization
	conf := dotenv.Init()
	err := conf.LoadConfig()
	util.ErrorHandler(errConf, err)

	// orm initialization
	orm, err := gorm.Open(
		postgres.Open(database.GetDSN(conf)),
		&gorm.Config{},
	)
	util.ErrorHandler(errDB, err)

	// run server
	defer glog.Flush()
	app.Run(conf, &database.Database{ORM: orm})
}
