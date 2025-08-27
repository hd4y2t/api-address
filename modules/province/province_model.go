package province

type Province struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"varchar(255)`
}
