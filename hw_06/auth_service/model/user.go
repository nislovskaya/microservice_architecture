package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"-" gorm:"not null"`
}
