package main

import (
	"log"
	"net/http"

	"github.com/Neiva07/review-digest-creator/api"
	"github.com/Neiva07/review-digest-creator/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	godotenv.Load()

	router.HandleFunc("/api/applicationDigests/{applicationId}", api.CreateApplicationDigest).Methods("POST")
	router.HandleFunc("/api/applicationDigests/{applicationId}", api.RemoveApplicationDigest).Methods("DELETE")

	services.ExecuteRunningJobs()

	log.Fatal(http.ListenAndServe(":8000", router))

}
