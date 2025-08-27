package city

type City struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	ProvinceID int    `json:"province_id" gorm:"not null"`
	Name       string `json:"name" gorm:"varchar(255)`
}
