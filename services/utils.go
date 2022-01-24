package services

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetJson(url string, target interface{}) error {
	log.Printf("Querying for %s url", url)

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	log.Println(r.Status)

	return json.NewDecoder(r.Body).Decode(target)
}
