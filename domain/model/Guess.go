package domain

type Guess struct {
	Id            int  `gorm:"type:int;primary_key;auto_increment"`
	MatchId       int  `gorm:"type:int" valid:"notnull"`
	UserId        int  `gorm:"type:int" valid:"notnull"`
	HomeTeamScore uint `gorm:"type:int;default:null"`
	AwayTeamScore uint `gorm:"type:int;default:null"`
	Points        uint `gorm:"type:int;default:null"`
}
