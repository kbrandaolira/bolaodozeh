package handler

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/useCase"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct {
	createUserUseCase useCase.CreateUserUseCase
}

func New(createUserUseCase useCase.CreateUserUseCase) handler {
	return handler{createUserUseCase: createUserUseCase}
}

func (h handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var user domain.User
	json.Unmarshal(body, &user)
	h.createUserUseCase.Execute(&user)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
