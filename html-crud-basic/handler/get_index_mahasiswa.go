package handler

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/ahmadzakyrifin/golang-http-projects/model"
)

type Data struct {
	Mahasiswa []model.Mahasiswa
}

var funcMap = template.FuncMap{
	"add" : func (a,b int)int{
			return a + b
	},
}

var indexTemplate = template.Must(
	template.New("index.html").Funcs(funcMap).ParseFiles(filepath.Join("views","index.html")),
)

func IndexMahasiswaHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, nim, address FROM mahasiswa ")
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}

		defer rows.Close()

		var mahasiswa []model.Mahasiswa
		for rows.Next(){
			var mhs model.Mahasiswa
			if err := rows.Scan(&mhs.ID,&mhs.Name,&mhs.Nim,&mhs.Address); err != nil {
				http.Error(w,err.Error(),http.StatusInternalServerError)
				return 
			}
			mahasiswa = append(mahasiswa, mhs)
		}

		data := Data{
			Mahasiswa: mahasiswa,
		}

		err = indexTemplate.Execute(w,data)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}

	}
}