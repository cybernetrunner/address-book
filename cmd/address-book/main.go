package main

import (
	log "github.com/sirupsen/logrus"
)

const (
	StrStart = "Service started"
	StrStop  = "Service stopped"
)

func main() {
	log.Info(StrStart)

	log.Info(StrStop)
}
