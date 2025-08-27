package district

type District struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	ProvinceId   string `json:"province_id" gorm:"column:province_id"`
	CityId      string `json:"city_id" gorm:"column:city_id"`
	Code        string `json:"code" gorm:"column:code"`
	Name        string `json:"name" gorm:"varchar(255)"`
}
