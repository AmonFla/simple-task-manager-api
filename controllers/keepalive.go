package controllers

import "net/http"

func KeepAlive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
