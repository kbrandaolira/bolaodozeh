package main

import (
	"bolaodozeh/application/handler"
	"bolaodozeh/domain/useCase"
	"bolaodozeh/infrastructure"
	"bolaodozeh/infrastructure/adapter"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := infrastructure.InitDB()

	userGoOrmAdapter := adapter.NewUserGoOrmAdapter(db)
	guessGoOrmAdapter := adapter.NewGuessGoOrmAdapter(db)
	matchGoOrmAdapter := adapter.NewMatchGoOrmAdapter(db)

	createBlankGuessForUserUseCase := useCase.NewCreateBlankGuessForUserUseCase(guessGoOrmAdapter, matchGoOrmAdapter)
	createUserUseCase := useCase.NewCreateUserUseCase(userGoOrmAdapter, createBlankGuessForUserUseCase)
	updateGuessUseCase := useCase.NewUpdateGuessUseCase(guessGoOrmAdapter, matchGoOrmAdapter)
	loginUseCase := useCase.NewLoginUseCase(userGoOrmAdapter)
	findGuessUseCase := useCase.NewFindGuessUseCase(guessGoOrmAdapter, matchGoOrmAdapter)
	updateMatchUseCase := useCase.NewUpdateMatchUseCase(matchGoOrmAdapter)
	findMatchUseCase := useCase.NewFindMatchUseCase(matchGoOrmAdapter)
	classificationUseCase := useCase.NewClassificationUseCase(guessGoOrmAdapter, matchGoOrmAdapter, userGoOrmAdapter)

	handler := handler.New(
		createUserUseCase,
		updateGuessUseCase,
		loginUseCase,
		findGuessUseCase,
		updateMatchUseCase,
		findMatchUseCase,
		classificationUseCase,
	)

	router := mux.NewRouter()
	router.HandleFunc("/user", handler.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/guess/{id}", handler.UpdateGuessHandler).Methods(http.MethodPut)
	router.HandleFunc("/guess/{userId}", handler.FindGuessHandler).Methods(http.MethodGet)
	router.HandleFunc("/match/{id}", handler.UpdateMatchHandler).Methods(http.MethodPut)
	router.HandleFunc("/match", handler.FindMatchHandler).Methods(http.MethodGet)
	router.HandleFunc("/classification", handler.ClassificationHandler).Methods(http.MethodGet)

	http.ListenAndServe(":4000", router)
}
