package useCase

import (
	"bolaodozeh/application/handler/dto"
	"bolaodozeh/domain/repository"

	"go.openly.dev/pointy"
)

type ClassificationUseCase struct {
	guessRepository repository.GuessRepository
	matchRepository repository.MatchRepository
	userRepository  repository.UserRepository
}

func NewClassificationUseCase(guessRepository repository.GuessRepository, matchRepository repository.MatchRepository, userRepository repository.UserRepository) ClassificationUseCase {
	return ClassificationUseCase{guessRepository: guessRepository, matchRepository: matchRepository, userRepository: userRepository}
}

func (c ClassificationUseCase) Execute() []dto.ClassificationDto {
	var dtos []dto.ClassificationDto
	users := c.userRepository.FindAll()
	for i := 0; i < len(users); i++ {
		var dto = dto.ClassificationDto{
			UserId:    users[i].Id,
			UserName:  users[i].Name,
			PlusThree: 0,
			PlusTwo:   0,
			PlusOne:   0,
			Points:    0,
		}
		matches := c.matchRepository.FindAll()
		for i := 0; i < len(matches); i++ {
			matchWithScore := matches[i].HomeTeamScore != nil && matches[i].AwayTeamScore != nil
			if matchWithScore {
				guess := c.guessRepository.FindByMatchIdAndUserId(matches[i].Id, users[i].Id)
				guessWithScore := guess.HomeTeamScore != nil && guess.AwayTeamScore != nil
				if guessWithScore {
					matchHomeWins := pointy.PointerValue(matches[i].HomeTeamScore, 0) > pointy.PointerValue(matches[i].AwayTeamScore, 0)
					matchAwayWins := pointy.PointerValue(matches[i].HomeTeamScore, 0) < pointy.PointerValue(matches[i].AwayTeamScore, 0)
					matchDraw := pointy.PointerValue(matches[i].HomeTeamScore, 0) == pointy.PointerValue(matches[i].AwayTeamScore, 0)

					guessHomeWins := pointy.PointerValue(guess.HomeTeamScore, 0) > pointy.PointerValue(guess.AwayTeamScore, 0)
					guessAwayWins := pointy.PointerValue(guess.HomeTeamScore, 0) < pointy.PointerValue(guess.AwayTeamScore, 0)
					guessDraw := pointy.PointerValue(guess.HomeTeamScore, 0) == pointy.PointerValue(guess.AwayTeamScore, 0)

					if (matchHomeWins && guessHomeWins) || (matchAwayWins && guessAwayWins) || (matchDraw && guessDraw) {
						plusThree := &matches[i].HomeTeamScore == &guess.HomeTeamScore && &matches[i].AwayTeamScore == &guess.AwayTeamScore
						plusOne := &matches[i].HomeTeamScore != &guess.HomeTeamScore && &matches[i].AwayTeamScore != &guess.AwayTeamScore
						if plusThree {
							dto.Points = dto.Points + 3
							dto.PlusThree++
						} else if plusOne {
							dto.Points = dto.Points + 1
							dto.PlusOne++
						} else {
							dto.Points = dto.Points + 2
							dto.PlusTwo++
						}
					}
				}
			}
		}

		dtos = append(dtos, dto)
	}
	return dtos
}
