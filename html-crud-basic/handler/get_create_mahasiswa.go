package handler

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

var createMahasiswaTemplate = template.Must(template.New("create.html").ParseFiles(filepath.Join("views","create.html")))

func GetCreateMahasiswaHandler(db *sql.DB) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := createMahasiswaTemplate.Execute(w,nil)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}
	}
}