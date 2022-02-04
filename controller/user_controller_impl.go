package controller

import (
	"encoding/json"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/model/web"
	"go-rest-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	userCreateRequest := web.UserCreateRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userCreateRequest)
	helper.PanicIfErr(err)

	userResponse := controller.UserService.CreateUser(r.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Status:  "200",
		Message: "OK",
		Data:    userResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfErr(err)
}

func (controller *UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUpdateRequest := web.UserUpdateRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userUpdateRequest)
	helper.PanicIfErr(err)

	params := mux.Vars(r)

	userUpdateRequest.Username = params["username"]

	userResponse := controller.UserService.UpdateUser(r.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Status:  "200",
		Message: "OK",
		Data:    userResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfErr(err)
}

func (controller *UserControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	controller.UserService.DeleteUser(r.Context(), username)
	webResponse := web.WebResponse{
		Status:  "200",
		Message: "OK",
		Data:    "User succesfully deleted",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfErr(err)
}

func (controller *UserControllerImpl) FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	userResponse := controller.UserService.FindUser(r.Context(), username)

	validToken, err := middleware.GenerateJWT(userResponse, w, r)
	helper.PanicIfErr(err)

	webResponse := web.WebResponse{
		Status:  "200",
		Message: validToken,
		Data:    userResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfErr(err)
}

func (controller *UserControllerImpl) FindAllUser(w http.ResponseWriter, r *http.Request) {
	userResponses := controller.UserService.FindAllUser(r.Context())

	webResponse := web.WebResponse{
		Status:  "200",
		Message: "OK",
		Data:    userResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfErr(err)
}
