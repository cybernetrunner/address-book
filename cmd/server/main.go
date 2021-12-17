package main

import (
	"github.com/cyneruxyz/address-book/internal/app"
	"github.com/cyneruxyz/address-book/internal/db"
	"github.com/cyneruxyz/address-book/pkg/util"
	"github.com/golang/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const errDB = "Fatal error db: %s \n"

func main() {
	// orm initialization
	orm, err := gorm.Open(
		postgres.Open(db.DSN),
		&gorm.Config{},
	)
	util.ErrorHandler(errDB, err)

	// run server
	defer glog.Flush()
	app.Run(&db.Database{ORM: orm})
}
