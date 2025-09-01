package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jywebs/go-rest-example/internal/handler"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/jywebs/go-rest-example/api"
)

// @title BusyBox Command Executor API
// @version 1.0
// @description This is a REST API server for executing commands in a BusyBox container
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := mux.NewRouter()
	
	// API v1 router
	api := r.PathPrefix("/api/v1").Subrouter()
	
	// Routes
	api.HandleFunc("/execute", handler.ExecuteCommand).Methods(http.MethodPost)
	
	// Swagger documentation
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	
	log.Printf("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
