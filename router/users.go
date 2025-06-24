package router

import (
	"minitask2/controller"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {
	r.GET("", controller.GetAlluser)
	r.GET("/:id", controller.GetOne)
	r.POST("", controller.CreateUser)
	r.PATCH("/:id", controller.UpdateUserCTRL)

}
