package controllers

import (
	"database/sql"
	"net/http"

	"github.com/carlos0934/videotube/auth"
)

type Middleware struct {
	conn *sql.DB
}

func (middleware *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	authUser := auth.NewUserAuth()
	authUser.Storage.GetConnection(middleware.conn)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if authUser.VefifyUser(token, &auth.UserClaims{}) {

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
