package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/dto"
	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/service"
)

type MahasiswaHandler struct {
	service service.MahasiswaService
}

func NewHandler (s service.MahasiswaService) *MahasiswaHandler {
	return &MahasiswaHandler{service : s}
}

func (h *MahasiswaHandler) Index  (w http.ResponseWriter, r *http.Request) {
	data,err := h.service.FindAll()
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
} 

func (h *MahasiswaHandler) Create (w http.ResponseWriter, r *http.Request) {
	var req dto.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	data,err := h.service.Create(req)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *MahasiswaHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req dto.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := h.service.Update(id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}


func (h *MahasiswaHandler) Delete (w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)	
}