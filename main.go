// Package main API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"
)

// swagger:route GET /home home
//
// Home endpoint.
//
// This will show all available endpoints.
//
//	Responses:
//	  default: genericError
//	  200: someResponse
func main() {
	log.Print("Starting the service...")

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", nil))

	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello! Your request was processed.")
	},
	)
	http.ListenAndServe(":8000", nil)
}
