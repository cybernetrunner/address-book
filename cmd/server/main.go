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

var (
	conf = &dotenv.DotEnv{
		ConfigFile: dotenv.DefaultConfigFile,
	}
	db  *database.Database
	orm *gorm.DB
	err error

	errConfig   = "Fatal error loading .env file: %s \n"
	errDatabase = "Fatal error database: %s \n"
	errServer   = "Fatal error server: %s \n"
)

func init() {
	// util initialization
	err = conf.LoadConfig()
	util.ErrorHandler(errConfig, err)

	// orm initialization
	orm, err = gorm.Open(
		postgres.Open(database.GetDSN(conf)),
		&gorm.Config{},
	)
	util.ErrorHandler(errDatabase, err)

	db = &database.Database{ORM: orm}
}

func main() {
	util.ErrorHandler(errServer, app.Run(conf, db))
	defer glog.Flush()
}
