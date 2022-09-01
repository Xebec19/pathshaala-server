package auth

import "net/http"

func Register(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Registration page working!"))
}