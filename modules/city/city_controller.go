package city

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CityController struct {
	cityService CityService
	ctx         *gin.Context
}

func NewCityController(cityService CityService, ctx *gin.Context) *CityController {
	return &CityController{cityService, ctx}
}

func (cc *CityController) GetAll() {
	cities := cc.cityService.GetAll()
	cc.ctx.JSON(200, cities)
}

func (cc *CityController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := cc.cityService.GetById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get city success",
		"data":    data,
	})
}

func (cc *CityController) GetByProvinceID(ctx *gin.Context) {
	provinceID, _ := strconv.Atoi(ctx.Param("province_id"))

	data := cc.cityService.GetByProvinceID(provinceID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get cities by province success",
		"data":    data,
	})
}
