package http

import (
	u "address-book/internal/usecase"

	"github.com/gin-gonic/gin"

	"net/http"
)

const (
	APIPath = "/address-field"
)

func SetupRouter() *gin.Engine {
	h := new(u.Handler)
	r := gin.Default()

	r.GET(
		"/ping",
		func(c *gin.Context) {
			c.String(
				http.StatusOK,
				"pong",
			)
		})

	r.POST(
		APIPath,
		func(c *gin.Context) {
			h.CreateField(c)
		},
	)

	r.GET(
		APIPath,
		func(c *gin.Context) {
			h.ReadField(c)
		},
	)

	r.PUT(
		APIPath,
		func(c *gin.Context) {
			h.UpdateField(c)
		},
	)

	r.DELETE(
		APIPath,
		func(c *gin.Context) {
			h.DeleteField(c)
		},
	)
	return r
}
