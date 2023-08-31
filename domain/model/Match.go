package domain

import (
	"time"
)

type Match struct {
	Id            int     `gorm:"type:int;primary_key;auto_increment"`
	HomeTeamName  string  `gorm:"type:varchar(100)" valid:"notnull"`
	HomeTeamScore *uint   `gorm:"type:int;default:null"`
	AwayTeamName  *string `gorm:"type:varchar(100)" valid:"notnull"`
	AwayTeamScore *uint   `gorm:"type:int;default:null"`
	DateTime      time.Time
	Stadium       string `gorm:"type:varchar(100)" valid:"notnull"`
	Phase         string `gorm:"type:varchar(100)" valid:"notnull"`
}
