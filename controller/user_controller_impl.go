package controller

import (
	"encoding/json"
	"fmt"
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

	// params := mux.Vars(r)
	// userUpdateRequest.Username = params["username"]

	username := middleware.GetCookie(w, r) // Get username by token in cookie
	strUsername := fmt.Sprint(username)
	if strUsername == "" {
		fmt.Fprint(w, "Please login first")
		return
	}
	userUpdateRequest.Username = strUsername

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
	// params := mux.Vars(r)
	// username := params["username"]

	username := middleware.GetCookie(w, r) // Get username by token in cookie
	strUsername := fmt.Sprint(username)
	if strUsername == "" {
		fmt.Fprint(w, "Please login first")
		return
	}

	controller.UserService.DeleteUser(r.Context(), strUsername)
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

	webResponse := web.WebResponse{
		Status:  "200",
		Message: "OK",
		Data:    userResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
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

func (controller *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userLoginRequest := web.UserLoginRequest{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userLoginRequest)
	helper.PanicIfErr(err)

	userResponse := controller.UserService.LoginUser(r.Context(), userLoginRequest)

	middleware.GenerateJWT(userResponse, w, r) // Generate token into token

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
