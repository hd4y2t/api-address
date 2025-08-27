package city

type City struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Code         string `json:"code" gorm:"column:code"`
	ProvinceCode string `json:"province_code" gorm:"column:province_code"`
	Name         string `json:"name" gorm:"varchar(255)`
}
