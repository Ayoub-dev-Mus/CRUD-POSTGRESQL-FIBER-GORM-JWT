package models

type User struct {
	Id         int       `gorm:"primary_key, AUTO_INCREMENT"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Comment    []Comment `gorm:"ForeignKey:UserId"`
	Order      []Order   `gorm:"ForeignKey:UserId"`
}
