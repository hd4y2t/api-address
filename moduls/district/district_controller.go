package district

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DistrictController struct {
	districtService DistrictService
	ctx             *gin.Context
}

func NewDistrictController(districtService DistrictService, ctx *gin.Context) *DistrictController {
	return &DistrictController{districtService, ctx}
}

func (dc *DistrictController) GetAll() {
	districts := dc.districtService.GetAll()
	dc.ctx.JSON(200, districts)
}

func (dc *DistrictController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := dc.districtService.GetById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get district success",
		"data":    data,
	})
}

func (dc *DistrictController) GetByProvinceID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := dc.districtService.GetByProvinceID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get district by province id success",
		"data":    data,
	})
}


func (dc *DistrictController) GetByCityID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := dc.districtService.GetByCityID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get district by city id success",
		"data":    data,
	})
}