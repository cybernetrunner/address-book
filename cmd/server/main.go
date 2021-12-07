package main

import (
	"github.com/cyneruxyz/address-book/internal/addressbook"
	"github.com/golang/glog"
)

func main() {
	if err := addressbook.Run(); err != nil {
		glog.Fatal(err)
	}

	defer glog.Flush()
}
