package handler

import (
	"bolaodozeh/application/handler/dto"
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/useCase"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handler struct {
	createUserUseCase  useCase.CreateUserUseCase
	updateGuessUseCase useCase.UpdateGuessUseCase
}

func New(createUserUseCase useCase.CreateUserUseCase, updateGuessUseCase useCase.UpdateGuessUseCase) handler {
	return handler{createUserUseCase: createUserUseCase, updateGuessUseCase: updateGuessUseCase}
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

func (h handler) UpdateGuessHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	params := mux.Vars(r)
	var dto dto.UpdateGuessDto
	json.Unmarshal(body, &dto)
	dto.GuessId, _ = strconv.Atoi(params["id"])
	h.updateGuessUseCase.Execute(dto)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Updated")
}
