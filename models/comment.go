package models



type Comment struct {
	Id        int `gorm:"primary_key, AUTO_INCREMENT"`
	Name      string
	Address   string
	FacultyID int `gorm:"column:faculty_id"`
	Faculty   User
}
