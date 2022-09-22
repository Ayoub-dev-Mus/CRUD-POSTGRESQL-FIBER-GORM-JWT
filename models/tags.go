package models

type Tags struct {
	TagId    int    `json:"id" gorm:"primaryKey"`
	Libeller string `json:"first_name"`
	GigId int `gorm:"column:Gig_id"`
	Gig  Gig
}
