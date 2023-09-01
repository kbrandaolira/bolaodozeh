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
	loginUseCase       useCase.LoginUseCase
	findGuessUseCase   useCase.FindGuessUseCase
}

func New(createUserUseCase useCase.CreateUserUseCase,
	updateGuessUseCase useCase.UpdateGuessUseCase,
	loginUseCase useCase.LoginUseCase,
	findGuessUseCase useCase.FindGuessUseCase,
) handler {
	return handler{
		createUserUseCase:  createUserUseCase,
		updateGuessUseCase: updateGuessUseCase,
		loginUseCase:       loginUseCase,
		findGuessUseCase:   findGuessUseCase,
	}
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
}

func (h handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var dto dto.LoginDto
	json.Unmarshal(body, &dto)
	logged := h.loginUseCase.Execute(&dto)
	w.Header().Add("Content-Type", "application/json")
	if logged {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
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
	w.WriteHeader(http.StatusNoContent)
}

func (h handler) FindGuessHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["userId"])

	guessDto, _ := json.Marshal(h.findGuessUseCase.Execute(userId))
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(guessDto)
}
