package handler

import (
	"database/sql"

	"net/http"
	"strconv"
)

func PostUpdateMahasiswaHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID tidak valid", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		nim := r.FormValue("nim")
		address := r.FormValue("address")

		_, err = db.Exec(
			"UPDATE mahasiswa SET name=?, nim=?, address=? WHERE id=?",
			name, nim, address, id,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/mahasiswa", http.StatusSeeOther)
	}
}
