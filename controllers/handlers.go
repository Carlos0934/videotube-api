package controllers

import (
	"fmt"
	"net/http"
)

func checkErr(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write(NewResponseMessage("Server error", true))
	}
}
