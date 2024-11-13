package route

import (
	"go-reggie/api/controller"
	"go-reggie/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	uc := &controller.HelloController{
		Env: env,
	}
	group.GET("/hello", uc.Hello)
}
