package main

import (
	"flag"
	"github.com/cyneruxyz/address-book/internal/addressbook"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := addressbook.Run(); err != nil {
		glog.Fatal(err)
	}
}
