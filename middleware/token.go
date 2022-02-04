package middleware

import (
	"fmt"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.RequestURI, "/api/user/") {
			next.ServeHTTP(w, r)
			return
		}

		if r.RequestURI == "/api/users" {
			cookie, err := r.Cookie("token")
			if err != nil {
				fmt.Fprintf(w, "No Cookie")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			tokenString := cookie.Value

			claims := jwt.MapClaims{}

			tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})
			helper.PanicIfErr(err)

			if !tkn.Valid {
				w.WriteHeader(http.StatusUnauthorized)
			}

			if claims["group_user"] != "admin" {
				fmt.Fprint(w, "Administratior required")
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func GenerateJWT(user web.UserResponse, w http.ResponseWriter, r *http.Request) (string, error) {

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["username"] = user.Username
	claims["password"] = user.Password
	claims["group_user"] = user.GroupUser
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	helper.PanicIfErr(err)

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Path = "/"
	http.SetCookie(w, cookie)

	return tokenString, nil
}
