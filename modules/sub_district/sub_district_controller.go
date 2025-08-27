package sub_district

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubDistrictController struct {
	subDistrictService SubDistrictService
	ctx                *gin.Context
}

func NewSubDistrictController(subDistrictService SubDistrictService, ctx *gin.Context) *SubDistrictController {
	return &SubDistrictController{subDistrictService, ctx}
}

func (dc *SubDistrictController) GetAll() {
	subDistricts := dc.subDistrictService.GetAll()
	dc.ctx.JSON(200, subDistricts)
}

func (dc *SubDistrictController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := dc.subDistrictService.GetById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get sub district success",
		"data":    data,
	})
}

func (dc *SubDistrictController) GetByProvinceID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := dc.subDistrictService.GetByProvinceID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get sub district by province id success",
		"data":    data,
	})
}


func (dc *SubDistrictController) GetByCityID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := dc.subDistrictService.GetByCityID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get sub district by city id success",
		"data":    data,
	})
}