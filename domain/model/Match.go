package domain

import (
	"time"
)

type Match struct {
	Id            int       `gorm:"type:int;primary_key;auto_increment"`
	HomeTeamName  string    `gorm:"type:varchar(100)" valid:"notnull"`
	HomeTeamScore *int      `gorm:"type:int;default:null"`
	AwayTeamName  string    `gorm:"type:varchar(100)" valid:"notnull"`
	AwayTeamScore *int      `gorm:"type:int;default:null"`
	DateTime      time.Time `gorm:"type:timestamp"`
	Stadium       string    `gorm:"type:varchar(100)" valid:"notnull"`
	Phase         string    `gorm:"type:varchar(100)" valid:"notnull"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:now()"`
}

type MatchStatus string

const (
	FINISHED     MatchStatus = ""
	NOT_FINISHED MatchStatus = ""
	Approved     MatchStatus = ""
	Rejected     MatchStatus = ""
)
