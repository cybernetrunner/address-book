package main

import (
	app "address-book/internal/address-book"
	log "github.com/sirupsen/logrus"
)

const (
	StrStop  = "[SERVER  STOP]"
	StrStart = "[SERVER START]"
)

func main() {
	defer log.Fatal(StrStop)

	log.Println(StrStart)
	app.Run()
}
