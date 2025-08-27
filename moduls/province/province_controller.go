package province

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProvinceController struct {
	provinceService ProvinceService
	ctx             *gin.Context
}

func NewProvinceController(provinceService ProvinceService, ctx *gin.Context) *ProvinceController {
	return &ProvinceController{provinceService, ctx}
}

func (pc *ProvinceController) GetAll() {
	provinces := pc.provinceService.GetAll()
	pc.ctx.JSON(200, provinces)
}

func (pc *ProvinceController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := pc.provinceService.GetById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get province success",
		"data":    data,
	})
}
