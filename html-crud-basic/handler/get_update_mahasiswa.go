package handler

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

var updateMahasiswaTemplate = template.Must(template.New("update.html").ParseFiles(filepath.Join("views","update.html")))

func GetUpdateMahasiswaHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID tidak valid", http.StatusBadRequest)
			return
		}

		row := db.QueryRow(
			"SELECT name, nim, address FROM mahasiswa WHERE id = ?",
			id,
		)

		var name, nim, address string
		err = row.Scan(&name, &nim, &address)
		if err != nil {
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		data := map[string]any{
			"id":      id,
			"name":    name,
			"nim":     nim,
			"address": address,
		}

		updateMahasiswaTemplate.Execute(w, data)
	}
}
