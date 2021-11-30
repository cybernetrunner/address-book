package usecase

import (
	"address-book/internal/usecase/repository"

	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	NotFoundResponse = "Address by this param not found"
	StatusOKResponse = "OK"
)

var repo = repository.NewRepository()

type Handler struct{}

func (h Handler) CreateField(c *gin.Context) {
	repo.AddField(
		c.PostForm("name"),
		c.PostForm("address"),
		c.PostForm("phone"),
	)
	c.JSON(http.StatusOK, StatusOKResponse)
}

func (h Handler) ReadField(c *gin.Context) {
	if i, ok := repo.FindField(c.PostForm("param")); ok {
		c.JSON(http.StatusOK, repo.GetItem(i))
	}
	c.JSON(http.StatusNotFound, NotFoundResponse)
}

func (h Handler) UpdateField(c *gin.Context) {
	if i, ok := repo.FindField(c.PostForm("param")); ok {
		repo.DeleteField(i)
		repo.AddField(
			c.PostForm("name"),
			c.PostForm("address"),
			c.PostForm("phone"),
		)

		c.JSON(http.StatusOK, StatusOKResponse)
	}
	c.JSON(http.StatusNotFound, NotFoundResponse)
}

func (h Handler) DeleteField(c *gin.Context) {
	if i, ok := repo.FindField(c.PostForm("param")); ok {
		repo.DeleteField(i)
		c.JSON(http.StatusOK, StatusOKResponse)
	}
	c.JSON(http.StatusNotFound, NotFoundResponse)
}
