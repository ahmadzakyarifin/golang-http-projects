package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func PostDeleteMahasiswaHandler(db *sql.DB) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id,err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return 
		}
		result,err := db.Exec("DELETE FROM mahasiswa WHERE id = ?",id)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}
		rowsAffected,err := result.RowsAffected()
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}
		fmt.Println("Row Delete : ",rowsAffected)
		http.Redirect(w, r, "/mahasiswa", http.StatusSeeOther)
	}
}