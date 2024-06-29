package main

import (
	"kyc/app/handlers"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// create server
func main() {
	addr := ":8080"
	http.HandleFunc("POST /sign-up", handlers.SignUp)
	http.HandleFunc("POST /users/{user_id}/submit-kyc", handlers.KycSubmit)
	http.HandleFunc("POST /users/{user_id}/kyc/action", handlers.ActionKYC)
	log.Info("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
