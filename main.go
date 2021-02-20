/*
 * User API
 *
 * An API for managing system users
 *
 * API version: 1.0.3
 * Contact: joe@jjssoftware.co.uk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	api "github.com/jjssoftware/go-user-api/go"
)

func main() {
	log.Printf("Server started")

	UserAPIApiService := api.NewUserAPIApiService()
	UserAPIApiController := api.NewUserAPIApiController(UserAPIApiService)

	router := api.NewRouter(UserAPIApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
