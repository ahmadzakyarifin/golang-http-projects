package routes

import (
	"net/http"

	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/handler"
)


func MapRoutes(mux *http.ServeMux, h *handler.MahasiswaHandler) {
    mux.HandleFunc("GET /mahasiswa", h.Index)
    mux.HandleFunc("POST /mahasiswa/create", h.Create)
    mux.HandleFunc("PATCH /mahasiswa/update/{id}", h.Update)
    mux.HandleFunc("DELETE /mahasiswa/delete/{id}", h.Delete)
}
