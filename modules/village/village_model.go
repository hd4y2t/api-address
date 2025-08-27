package village

type Village struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	ProvinceCode string `json:"province_code" gorm:"column:province_code"`
	CityCode     string `json:"city_code" gorm:"column:city_code"`
	DistrictCode string `json:"district_code" gorm:"column:district_code"`
	Code         string `json:"code" gorm:"column:code"`
	Name         string `json:"name" gorm:"varchar(255)`
}
