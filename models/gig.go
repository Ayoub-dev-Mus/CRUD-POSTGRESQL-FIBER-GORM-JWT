package models

type Gig struct {
	GigId       int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Order    []Order `gorm:"ForeignKey:GigId"`
	Tag    []Tags `gorm:"ForeignKey:GigId"`


}
