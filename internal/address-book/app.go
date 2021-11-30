package address_book

import (
	controller "address-book/internal/routers/http"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	UriHost = "localhost"
	UriPort = "8000"
)

var (
	router  = controller.SetupRouter()
	address = fmt.Sprintf("%s:%s", UriHost, UriPort)
)

func Run() {
	if err := router.Run(address); err != nil {
		log.Errorln(err)
		return
	}
}
