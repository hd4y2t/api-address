package village

type Village struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	ProvinceId   string `json:"province_id" gorm:"column:province_id"`
	CityId      string `json:"city_id" gorm:"column:city_id"`
	DistrictId   string `json:"district_id" gorm:"column:district_id"`
	Code         string `json:"code" gorm:"column:code"`
	Name         string `json:"name" gorm:"varchar(255)`
}
