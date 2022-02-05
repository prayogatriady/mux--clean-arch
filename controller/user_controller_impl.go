package controller

import (
	"encoding/json"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/model/web"
	"go-rest-api/service"
	"log"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Signup(w http.ResponseWriter, r *http.Request) {

	var webResponse web.WebResponse

	tokenData := middleware.GetCookie(w, r)
	log.Println(tokenData)

	// check if user already login
	if tokenData.Username != "" {
		webResponse = web.WebResponse{
			Status:  "400",
			Message: "BAD REQUEST",
			Data:    "Already logged in, logout first to signing up",
		}
	} else {

		userCreateRequest := web.UserCreateRequest{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userCreateRequest)
		helper.PanicIfErr(err)

		userResponse := controller.UserService.FindUser(r.Context(), userCreateRequest.Username)

		// check if user exists
		if userResponse.Username != "" {
			webResponse = web.WebResponse{
				Status:  "400",
				Message: "BAD REQUEST",
				Data:    "User already exists",
			}
		} else {
			userResponse = controller.UserService.CreateUser(r.Context(), userCreateRequest)
			webResponse = web.WebResponse{
				Status:  "201",
				Message: "CREATED",
				Data:    userResponse,
			}

			middleware.GenerateJWT(userResponse, w, r) // Generate new user data into token
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var webResponse web.WebResponse

	tokenData := middleware.GetCookie(w, r)
	log.Println(tokenData)

	// check if user not logged in
	if tokenData.Username == "" {
		webResponse = web.WebResponse{
			Status:  "401",
			Message: "UNAUTHORIZED",
			Data:    "Not logged in",
		}
	} else {
		userUpdateRequest := web.UserUpdateRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userUpdateRequest)
		helper.PanicIfErr(err)

		userUpdateRequest.Username = tokenData.Username

		userResponse := controller.UserService.UpdateUser(r.Context(), userUpdateRequest)
		webResponse = web.WebResponse{
			Status:  "202",
			Message: "ACCEPTED",
			Data:    userResponse,
		}

		middleware.GenerateJWT(userResponse, w, r) // Generate new user data into token
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {

	var webResponse web.WebResponse

	tokenData := middleware.GetCookie(w, r)
	log.Println(tokenData)

	// check if user not logged in
	if tokenData.Username == "" {
		webResponse = web.WebResponse{
			Status:  "401",
			Message: "UNAUTHORIZED",
			Data:    "Not logged in",
		}
	} else {
		controller.UserService.DeleteUser(r.Context(), tokenData.Username)
		webResponse = web.WebResponse{
			Status:  "200",
			Message: "OK",
			Data:    "User succesfully deleted",
		}

		middleware.DeleteCookie(w, r) // delete cookie
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Profile(w http.ResponseWriter, r *http.Request) {

	var webResponse web.WebResponse

	tokenData := middleware.GetCookie(w, r)
	log.Println(tokenData)

	// check if user not logged in
	if tokenData.Username == "" {
		webResponse = web.WebResponse{
			Status:  "401",
			Message: "UNAUTHORIZED",
			Data:    "Not logged in",
		}
	} else {
		userResponse := controller.UserService.FindUser(r.Context(), tokenData.Username)

		webResponse = web.WebResponse{
			Status:  "200",
			Message: "OK",
			Data:    userResponse,
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) FindAllUser(w http.ResponseWriter, r *http.Request) {

	var webResponse web.WebResponse

	tokenData := middleware.GetCookie(w, r)
	log.Println(tokenData)

	// check if user not logged in
	if tokenData.Username == "" {
		webResponse = web.WebResponse{
			Status:  "401",
			Message: "UNAUTHORIZED",
			Data:    "Not logged in",
		}
	} else {

		// Only admin can access this endpoint
		if tokenData.GroupUser == "admin" {
			userResponses := controller.UserService.FindAllUser(r.Context())

			webResponse = web.WebResponse{
				Status:  "200",
				Message: "OK",
				Data:    userResponses,
			}
		} else {
			webResponse = web.WebResponse{
				Status:  "405",
				Message: "METHOD NOT ALLOWED",
				Data:    "Admin required",
			}
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {

	var webResponse web.WebResponse

	tokenData := middleware.GetCookie(w, r)
	log.Println(tokenData)

	// check if user already login
	if tokenData.Username != "" {
		webResponse = web.WebResponse{
			Status:  "400",
			Message: "BAD REQUEST",
			Data:    "Already logged in",
		}
	} else {

		userLoginRequest := web.UserLoginRequest{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userLoginRequest)
		helper.PanicIfErr(err)

		userResponse := controller.UserService.LoginUser(r.Context(), userLoginRequest)

		middleware.GenerateJWT(userResponse, w, r) // Generate user data into token

		webResponse = web.WebResponse{
			Status:  "200",
			Message: "OK",
			Data:    userResponse,
		}
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request) {

	var webResponse web.WebResponse

	tokenData := middleware.GetCookie(w, r)
	log.Println(tokenData)

	// check if user not logged in
	if tokenData.Username == "" {
		webResponse = web.WebResponse{
			Status:  "401",
			Message: "UNAUTHORIZED",
			Data:    "Not logged in",
		}
	} else {

		webResponse = web.WebResponse{
			Status:  "200",
			Message: "OK",
			Data:    "Logout successful",
		}

		middleware.DeleteCookie(w, r) // delete cookie
	}

	helper.WriteToResponseBody(w, webResponse)
}
