package domain

type User struct {
	Id       int    `gorm:"type:int;primary_key;auto_increment"`
	Name     string `gorm:"type:varchar(255);not null" valid:"notnull"`
	Email    string `gorm:"type:varchar(255);unique;not null" valid:"notnull,email"`
	Password string `gorm:"type:varchar(255);not null" valid:"notnull"`
	Admin    bool   `gorm:"type:boolean;default:0;not null"`
}
