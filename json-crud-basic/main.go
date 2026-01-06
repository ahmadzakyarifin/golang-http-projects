package main

import (
	"fmt"
	"net/http"

	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/database"
	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/handler"
	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/repo"
	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/routes"
	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/service"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	mahasiswaRepo := repo.NewMahasiswaRepo(db)
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepo)
	handler := handler.NewHandler(mahasiswaService)

	mux := http.NewServeMux()
	routes.MapRoutes(mux,handler)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}