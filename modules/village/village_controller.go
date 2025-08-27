package village

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VillageController struct {
	villageService VillageService
	ctx            *gin.Context
}

func NewVillageController(villageService VillageService, ctx *gin.Context) *VillageController {
	return &VillageController{villageService, ctx}
}

func (vc *VillageController) GetAll() {
	villages := vc.villageService.GetAll()
	vc.ctx.JSON(200, villages)
}

func (vc *VillageController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := vc.villageService.GetById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get village success",
		"data":    data,
	})
}

func (vc *VillageController) GetByProvinceID(ctx *gin.Context) {
	provinceID, _ := strconv.Atoi(ctx.Param("province_id"))

	data := vc.villageService.GetByProvinceID(provinceID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get villages by province success",
		"data":    data,
	})
}

func (vc *VillageController) GetByCityID(ctx *gin.Context) {
	cityID, _ := strconv.Atoi(ctx.Param("city_id"))

	data := vc.villageService.GetByCityID(cityID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get villages by city success",
		"data":    data,
	})
}

func (vc *VillageController) GetBySubDistrictID(ctx *gin.Context) {
	subDistrictID, _ := strconv.Atoi(ctx.Param("sub_district_id"))

	data := vc.villageService.GetBySubDistrictID(subDistrictID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get villages by sub-district success",
		"data":    data,
	})
}
