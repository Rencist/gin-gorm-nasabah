package controller

import (
	"gin-gorm-nasabah/common"
	"gin-gorm-nasabah/dto"
	"gin-gorm-nasabah/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RekeningController interface {
	CreateRekening(ctx *gin.Context)
	DeleteRekening(ctx *gin.Context)
}

type rekeningController struct {
	rekeningService service.RekeningService
}

func NewRekeningController(rs service.RekeningService) RekeningController {
	return &rekeningController{
		rekeningService: rs,
	}
}

func(rc *rekeningController) CreateRekening(ctx *gin.Context) {
	var rekening dto.RekeningCreateDto
	err := ctx.ShouldBind(&rekening)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Rekening", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := rc.rekeningService.CreateRekening(ctx.Request.Context(), rekening)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Rekening", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menambahkan Rekening", result)
	ctx.JSON(http.StatusOK, res)
}

func(rc *rekeningController) DeleteRekening(ctx *gin.Context) {
	rekeningID := ctx.Param("id")
	_, err := rc.rekeningService.DeleteRekening(ctx.Request.Context(), rekeningID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Rekening", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Rekening", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}