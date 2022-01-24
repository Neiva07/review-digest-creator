package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Neiva07/review-digest-creator/persistence"
	"github.com/Neiva07/review-digest-creator/services"
	"github.com/gorilla/mux"
)

type CreateApplicationDigestRequest struct {
	StartDigestAt time.Time `json:"startDigestAt"`
	SleepTime     uint64    `json:"sleepTime"`
	PageSize      uint64    `json:"pageSize"`
}

func CreateApplicationDigest(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	applicationId := params["applicationId"]

	log.Printf("Creating Application Digest for application %s", applicationId)

	request := &CreateApplicationDigestRequest{
		SleepTime:     24,
		StartDigestAt: time.Now(),
		PageSize:      10,
	}

	json.NewDecoder(r.Body).Decode(request)

	application := persistence.AddApplicationAndActivate(applicationId, request.SleepTime, request.StartDigestAt, request.PageSize)

	services.ExecuteNewJob(application)
}

func RemoveApplicationDigest(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	applicationId := params["applicationId"]

	log.Printf("Deleting Application Digest for application %s", applicationId)

	persistence.InactivateApplicationDigest(applicationId)
}
