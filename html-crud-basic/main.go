package main

import (
	"fmt"
	"net/http"

	"github.com/ahmadzakyrifin/golang-http-projects/database"
	"github.com/ahmadzakyrifin/golang-http-projects/routes"
)

func main(){
	db := database.InitDB()

	mux := http.NewServeMux()

	routes.MapRoutes(mux,db)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	fmt.Println("http://localhost:8080")
	server.ListenAndServe()
}