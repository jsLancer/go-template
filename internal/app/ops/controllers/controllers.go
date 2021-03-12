package controllers

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/pkg/http"
)

func InitControllersFn(
	demoCtl *DemoController,
) http.InitControllers {
	return func(r *gin.Engine) {

		r.GET("/demo", demoCtl.FindAll)
		r.GET("/demo/:id", demoCtl.GetByID)
		r.POST("demo", demoCtl.Create)
		r.PUT("/demo/:id", demoCtl.Update)
		r.DELETE("/demo", demoCtl.DeleteByID)

	}
}
