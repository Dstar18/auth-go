package models

type User struct {
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
