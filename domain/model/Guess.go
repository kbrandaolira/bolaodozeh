package domain

import "time"

type Guess struct {
	Id            int       `gorm:"type:int;primary_key;auto_increment"`
	MatchId       int       `gorm:"type:int" valid:"notnull"`
	UserId        int       `gorm:"type:int" valid:"notnull"`
	HomeTeamScore *int      `gorm:"type:int;default:null"`
	AwayTeamScore *int      `gorm:"type:int;default:null"`
	Points        *int      `gorm:"type:int;default:null"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:now()"`
}
