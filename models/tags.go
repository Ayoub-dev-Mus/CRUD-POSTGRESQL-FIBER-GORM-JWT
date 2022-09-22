package models

type Tags struct {
	TagId    int    `json:"id" gorm:"primaryKey"`
	Libeller string `json:"first_name"`
}
