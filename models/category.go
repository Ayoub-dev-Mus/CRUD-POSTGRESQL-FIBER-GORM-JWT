package models

type Category struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Libeller string `json:"libeller"`
}
