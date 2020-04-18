package api

import (
	"os"
	"net/http"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
)

func SetupHTTPUpdates(apiToken string, updateFunction func()) error {
	if apiToken == "" {
		return errors.New("API token is empty or has not been set. Not starting API.")
	}
	
	log.Println("Watchtower HTTP API started.")

	http.HandleFunc("/v1/update", func(w http.ResponseWriter, r *http.Request){
		log.Info("Updates triggered by HTTP API request.")
		
		_, err := io.Copy(os.Stdout, r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		if r.Header.Get("Token") != apiToken {
			log.Println("Invalid token. Not updating.")
			return
		}

		log.Println("Valid token found. Triggering updates.")
		
		updateFunction()
	})
	
	return nil
}

func WaitForHTTPUpdates() error {
	log.Fatal(http.ListenAndServe(":8080", nil))
	os.Exit(0)
	return nil
}