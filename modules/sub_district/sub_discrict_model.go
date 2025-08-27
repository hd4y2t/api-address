package sub_district

type SubDistrict struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	ProvinceID int    `json:"province_id" gorm:"column:province_id"`
	CityId     int    `json:"city_id" gorm:"column:city_id"`
	Name       string `json:"name" gorm:"varchar(255)"`
}
