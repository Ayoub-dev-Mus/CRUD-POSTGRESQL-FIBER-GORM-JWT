package models

type Comment struct {
	Id        int `gorm:"primary_key, AUTO_INCREMENT"`
	Libeller     string `json:"Libeller"`
	UserId int `gorm:"column:User_id"`
	User  User
}