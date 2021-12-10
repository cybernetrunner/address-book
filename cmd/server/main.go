package main

import (
	"github.com/cyneruxyz/address-book/internal/app"
	"github.com/golang/glog"
)

func main() {
	if err := app.Run(); err != nil {
		glog.Fatal(err)
	}

	defer glog.Flush()
}
