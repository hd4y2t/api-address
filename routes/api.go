package routes

import (
	city "api-address/moduls/city"
	district "api-address/moduls/district"
	province "api-address/moduls/province"
	village "api-address/moduls/village"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB) {
	provinceRepository := province.NewProvinceRepository(db)
	provinceService := province.NewProvinceService(provinceRepository)
	provinceController := province.NewProvinceController(provinceService, ctx)

	cityRepository := city.NewCityRepository(db)
	cityService := city.NewCityService(cityRepository)
	cityController := city.NewCityController(cityService, ctx)

	districtRepository := district.NewDistrictRepository(db)
	districtService := district.NewDistrictService(districtRepository)
	districtController := district.NewDistrictController(districtService, ctx)

	villageRepository := village.NewVillageRepository(db)
	villageService := village.NewVillageService(villageRepository)
	villageController := village.NewVillageController(villageService, ctx)

	api := router.Group("/api")
	{
		// province
		api.GET("/province", provinceController.GetAll)
		api.GET("/province/:id", provinceController.GetById)

		//city
		api.GET("/city/:id", cityController.GetByProvinceID)

		//district
		api.GET("/district/:id", districtController.GetById)

		//village
		api.GET("/village/:id", villageController.GetById)
	}
}
