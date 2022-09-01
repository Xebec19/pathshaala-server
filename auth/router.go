package auth

import "github.com/gorilla/mux"

func Router(r *mux.Router) {
	routes := r.PathPrefix("/auth").Subrouter()

	routes.HandleFunc("/register", Register)
}
