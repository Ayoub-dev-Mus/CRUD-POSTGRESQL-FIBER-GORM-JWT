package models


type Order struct {
	OrderId    int `json:"oderid"`
	Order_date string `json:"order_date"`
	GigId     int `gorm:"column:Gig_id"`
	Gig       Gig
	UserId     int `gorm:"column:User_id"`
	User       User
}