package main

import (
	"net/http"

	"github.com/Xebec19/pathshaala-server/auth"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	auth.Router(r)
	http.ListenAndServe(":8080", r)
}
