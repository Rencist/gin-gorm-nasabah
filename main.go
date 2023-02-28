package main

import (
	"gin-gorm-nasabah/common"
	"gin-gorm-nasabah/config"
	"gin-gorm-nasabah/controller"
	"gin-gorm-nasabah/repository"
	"gin-gorm-nasabah/routes"
	"gin-gorm-nasabah/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		res := common.BuildErrorResponse("Gagal Terhubung ke Server", err.Error(), common.EmptyObj{})
		(*gin.Context).JSON((&gin.Context{}), http.StatusBadGateway, res)
		return
	}

	var (
		db *gorm.DB = config.SetupDatabaseConnection()
		
		nasabahRepository repository.NasabahRepository = repository.NewNasabahRepository(db)
		nasabahService service.NasabahService = service.NewNasabahService(nasabahRepository)
		nasabahController controller.NasabahController = controller.NewNasabahController(nasabahService)
		
		rekeningRepository repository.RekeningRepository = repository.NewRekeningRepository(db)
		rekeningService service.RekeningService = service.NewRekeningService(rekeningRepository)
		rekeningController controller.RekeningController = controller.NewRekeningController(rekeningService)
	)

	server := gin.Default()
	routes.NasabahRoutes(server, nasabahController)
	routes.RekeningRoutes(server, rekeningController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run("127.0.0.1:" + port)
}