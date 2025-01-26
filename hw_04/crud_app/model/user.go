package model

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string `json:"username" gorm:"type:varchar(100);unique;not null"`
	FirstName string `json:"firstName" gorm:"size:100;not null"`
	LastName  string `json:"lastName" gorm:"size:100;not null"`
	Email     string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone     string `json:"phone" gorm:"type:varchar(100);unique;not null"`
}
