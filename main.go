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
	//db
	db := infrastructure.InitDB()
	// app classes
	userGoOrmAdapter := adapter.NewUserGoOrmAdapter(db)
	guessGoOrmAdapter := adapter.NewGuessGoOrmAdapter(db)
	matchGoOrmAdapter := adapter.NewMatchGoOrmAdapter(db)
	createBlankGuessForUserUseCase := useCase.NewCreateBlankGuessForUserUseCase(guessGoOrmAdapter, matchGoOrmAdapter)
	createUserUseCase := useCase.NewCreateUserUseCase(userGoOrmAdapter, createBlankGuessForUserUseCase)
	handler := handler.New(createUserUseCase)
	// routes
	router := mux.NewRouter()
	router.HandleFunc("/users", handler.CreateUserHandler).Methods(http.MethodPost)
	// start
	http.ListenAndServe(":4000", router)
}
