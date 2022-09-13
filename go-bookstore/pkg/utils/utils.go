package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("an err occured reading request body: %v", err)
	}

	err = json.Unmarshal(body, &x)
	if err != nil {
		log.Fatalf("an err occured unmarchalling json: %v", err)
	}
}
