package routes

import (
	"database/sql"
	"net/http"

	"github.com/ahmadzakyrifin/golang-http-projects/handler"
)

func MapRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("GET /mahasiswa",handler.IndexMahasiswaHandler(db))
	mux.HandleFunc("GET /mahasiswa/create",handler.GetCreateMahasiswaHandler(db))
	mux.HandleFunc("POST /mahasiswa/create",handler.PostCreateMahasiswaHandler(db))
	mux.HandleFunc("GET /mahasiswa/update/{id}",handler.GetUpdateMahasiswaHandler(db))
	mux.HandleFunc("POST /mahasiswa/update/{id}",handler.PostUpdateMahasiswaHandler(db))
	mux.HandleFunc("POST /mahasiswa/delete/{id}",handler.PostDeleteMahasiswaHandler(db))
}
