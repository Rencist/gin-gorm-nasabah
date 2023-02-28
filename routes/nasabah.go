package routes

import (
	"gin-gorm-nasabah/controller"

	"github.com/gin-gonic/gin"
)

func NasabahRoutes(router *gin.Engine, NasabahController controller.NasabahController) {
	nasabahRoutes := router.Group("/api/nasabah")
	{
		nasabahRoutes.GET("", NasabahController.GetAllNasabah)
		nasabahRoutes.GET("/:id", NasabahController.GetNasabahByID)
		nasabahRoutes.POST("", NasabahController.CreateNasabah)
		nasabahRoutes.PUT("/:id", NasabahController.UpdateNasabah)
		nasabahRoutes.DELETE("/:id", NasabahController.DeleteNasabah)
	}
}