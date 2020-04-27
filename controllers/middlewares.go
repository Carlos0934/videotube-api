package controllers

import (
	"fmt"
	"net/http"

	"github.com/carlos0934/videotube/auth"
)

type UserMiddleware struct {
	userAuth auth.UserAuth
}

func NewUserMiddleware(userAuth *auth.UserAuth) *UserMiddleware {

	return &UserMiddleware{
		userAuth: *userAuth,
	}
}
func (middleware *UserMiddleware) AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		fmt.Println("dsadsa")
		if middleware.userAuth.VefifyUser(token, &auth.UserClaims{}) {

			next.ServeHTTP(w, r)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(NewResponseMessage("Invalid Token Try again", false))
		}

	})
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
