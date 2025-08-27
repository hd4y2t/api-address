package village

type Village struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	ProvinceID    int    `json:"province_id" gorm:"not null"`
	CityId        int    `json:"city_id" gorm:"column:city_id"`
	SubDistrictID int    `json:"sub_district_id" gorm:"column:sub_district_id"`
	Name          string `json:"name" gorm:"varchar(255)`
}
