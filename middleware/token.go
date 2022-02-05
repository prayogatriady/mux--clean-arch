package middleware

import (
	"context"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Method, "-", r.RequestURI)

		// get cookie with name token
		cookie, _ := r.Cookie("token")
		if cookie != nil {
			tokenString := cookie.Value

			claims := jwt.MapClaims{}

			token, _ := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})

			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
			}

			ctx := context.WithValue(r.Context(), "token", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func GenerateJWT(user web.UserResponse, w http.ResponseWriter, r *http.Request) {

	claims := jwt.MapClaims{}

	expTime := time.Now().Add(time.Minute * 1)

	claims["authorized"] = true
	claims["username"] = user.Username
	claims["password"] = user.Password
	claims["group_user"] = user.GroupUser
	claims["email"] = user.Email
	claims["exp"] = expTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	helper.PanicIfErr(err)

	// create cookie with token as a name
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = expTime
	http.SetCookie(w, cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request) web.UserResponse {
	var (
		claims       interface{}
		data         jwt.MapClaims
		userResponse web.UserResponse
	)

	if claims = r.Context().Value("token"); claims != nil {
		data = claims.(jwt.MapClaims)
		userResponse = web.UserResponse{
			Username:  data["username"].(string),
			Password:  data["password"].(string),
			GroupUser: data["group_user"].(string),
			Email:     data["email"].(string),
		}
		return userResponse
	} else {
		return userResponse
	}
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, -1) // Set to expire in the past
	cookie := http.Cookie{Name: "token", Value: "", Expires: expiration}
	http.SetCookie(w, &cookie)
}
