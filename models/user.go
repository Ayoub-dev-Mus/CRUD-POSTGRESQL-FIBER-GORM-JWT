package models

type User struct {
	Id       int `gorm:"primary_key, AUTO_INCREMENT"`
	Name     string
	Students []Comment `gorm:"ForeignKey:FacultyID"`

}
