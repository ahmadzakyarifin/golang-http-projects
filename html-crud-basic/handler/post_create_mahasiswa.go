package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)


func PostCreateMahasiswaHandler(db *sql.DB)func (w http.ResponseWriter, r *http.Request) {
	return  func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		nim := r.FormValue("nim")
		address := r.FormValue("address")

		result,err := db.Exec("INSERT INTO mahasiswa (name,nim,address) VALUES (?,?,?)",name,nim,address)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}
		rowsAffected,err := result.RowsAffected()
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}
		fmt.Println("Row Insert : ",rowsAffected)

		http.Redirect(w,r,"/mahasiswa",http.StatusSeeOther)
	}
}