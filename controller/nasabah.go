package controller

import (
	"gin-gorm-nasabah/common"
	"gin-gorm-nasabah/dto"
	"gin-gorm-nasabah/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NasabahController interface {
	GetAllNasabah(ctx *gin.Context)
	CreateNasabah(ctx *gin.Context)
}

type nasabahController struct {
	nasabahService service.NasabahService
}

func NewNasabahController(ns service.NasabahService) NasabahController {
	return &nasabahController{
		nasabahService: ns,
	}
}

func(nc *nasabahController) GetAllNasabah(ctx *gin.Context) {
	result, err := nc.nasabahService.GetAllNasabah(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Nasabah", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Nasabah", result)
	ctx.JSON(http.StatusOK, res)
}

func(nc *nasabahController) CreateNasabah(ctx *gin.Context) {
	var nasabah dto.NasabahCreateDto
	err := ctx.ShouldBind(&nasabah)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Nasabah", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := nc.nasabahService.CreateNasabah(ctx.Request.Context(), nasabah)
	if err!= nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Nasabah", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Nasabah", result)
	ctx.JSON(http.StatusOK, res)
}