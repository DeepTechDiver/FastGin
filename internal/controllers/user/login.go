package user

import (
	"fast-gin/internal/controllers"
	receive "fast-gin/internal/interaction/receive/user"
	"fast-gin/internal/service/user"
	"github.com/gin-gonic/gin"
)

type LoginControllers struct {
	controllers.BaseControllers
}

// Login 登入
func (lg LoginControllers) Login(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.LoginReceive)); err == nil {
		results, err := user.Login(rec)
		lg.Response(ctx, results, err)
	}
}

// Register 注册
func (lg LoginControllers) Register(ctx *gin.Context) {
	if rec, err := controllers.ShouldBind(ctx, new(receive.RegisterReceive)); err == nil {
		results, err := user.Register(rec)
		lg.Response(ctx, results, err)
	}
}
