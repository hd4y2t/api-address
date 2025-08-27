package province

type Province struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"varchar(10)"`
	Name string `json:"name" gorm:"varchar(255)"`
}
