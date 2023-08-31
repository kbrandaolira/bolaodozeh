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
	updateGuessUseCase := useCase.NewUpdateGuessUseCase(guessGoOrmAdapter, matchGoOrmAdapter)
	handler := handler.New(createUserUseCase, updateGuessUseCase)
	// routes
	router := mux.NewRouter()
	router.HandleFunc("/user", handler.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/guess/{id}", handler.UpdateGuessHandler).Methods(http.MethodPut)
	// start
	http.ListenAndServe(":4000", router)
}
