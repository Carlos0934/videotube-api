package controllers

import (
	"net/http"

	"github.com/carlos0934/videotube/auth"
	"github.com/gorilla/mux"
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
		if mux.Vars(r)["user"] == "" {

			next.ServeHTTP(w, r)
			return
		}
		token := r.Header.Get("Authorization")

		if middleware.userAuth.VefifyUser(token, &auth.UserClaims{}) {

			next.ServeHTTP(w, r)

		} else {
			w.Header().Add("Content-Type", "application/json")
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
