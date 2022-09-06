package main

import (
	"net/http"

	"github.com/emmanueltukpe/go-web-server/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":3000", srv)
}