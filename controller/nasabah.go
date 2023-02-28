package controller

import (
	"gin-gorm-nasabah/common"
	"gin-gorm-nasabah/dto"
	"gin-gorm-nasabah/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NasabahController interface {
	GetAllNasabah(ctx *gin.Context)
	CreateNasabah(ctx *gin.Context)
	UpdateNasabah(ctx *gin.Context)
	DeleteNasabah(ctx *gin.Context)
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

func(nc *nasabahController) UpdateNasabah(ctx *gin.Context) {
	nasabahID := ctx.Param("id")
	var nasabah dto.NasabahUpdateDto
	err := ctx.ShouldBind(&nasabah)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Nasabah", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	nasabah.ID, _ = uuid.Parse(nasabahID)
	result, err := nc.nasabahService.UpdateNasabah(ctx.Request.Context(), nasabah)
	if err!= nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Nasabah", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mengupdate Nasabah", result)
	ctx.JSON(http.StatusOK, res)
}

func(nc *nasabahController) DeleteNasabah(ctx *gin.Context) {
	nasabahID := ctx.Param("id")
	_, err := nc.nasabahService.DeleteNasabah(ctx.Request.Context(), nasabahID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Nasabah", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menghapus Nasabah", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}