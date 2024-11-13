package controller

import (
	"github.com/gin-gonic/gin"
	"go-reggie/bootstrap"
	"net/http"
)

type HelloController struct {
	Env *bootstrap.Env
}

func (hc *HelloController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})

}
