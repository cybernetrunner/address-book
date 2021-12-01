package http

import (
	u "address-book/internal/usecase"
	r "address-book/internal/usecase/repository"

	"github.com/gin-gonic/gin"

	"net/http"
)

const (
	APIPath = "/address-field"
)

func SetupRouter(repo *r.Repository) *gin.Engine {
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
			h.CreateField(c, repo)
		},
	)

	r.GET(
		APIPath,
		func(c *gin.Context) {
			h.ReadField(c, repo)
		},
	)

	r.PUT(
		APIPath,
		func(c *gin.Context) {
			h.UpdateField(c, repo)
		},
	)

	r.DELETE(
		APIPath,
		func(c *gin.Context) {
			h.DeleteField(c, repo)
		},
	)
	return r
}
