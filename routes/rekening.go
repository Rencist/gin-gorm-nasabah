package routes

import (
	"gin-gorm-nasabah/controller"

	"github.com/gin-gonic/gin"
)

func RekeningRoutes(router *gin.Engine, RekeningController controller.RekeningController) {
	rekeningRoutes := router.Group("/api/rekening")
	{
		rekeningRoutes.POST("", RekeningController.CreateRekening)
		rekeningRoutes.DELETE("/:id", RekeningController.DeleteRekening)
	}
}