package repo

import (
	"database/sql"
	"errors"

	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/model"
)

type MahasiswaRepository interface {
    FindAll() ([]model.Mahasiswa, error)
    Create(m model.Mahasiswa) (model.Mahasiswa, error)
    Update(m model.Mahasiswa) (model.Mahasiswa, error)
    Delete(id int) error
}

type mahasiswaRepo struct {
	db *sql.DB
}

func NewMahasiswaRepo(db *sql.DB) MahasiswaRepository {
	return &mahasiswaRepo{db: db}
}

func (r *mahasiswaRepo) FindAll() ([]model.Mahasiswa, error) {
	rows,err := r.db.Query("SELECT id,name,nim,address FROM mahasiswa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var mahasiswas []model.Mahasiswa
	for rows.Next() {
		var m model.Mahasiswa
		err := rows.Scan(&m.ID,&m.Name,&m.Nim,&m.Address)
		if err != nil {
			return nil, err
		}
		mahasiswas = append(mahasiswas, m)
	}
	return mahasiswas, nil
}

func (r *mahasiswaRepo) Create(m model.Mahasiswa) (model.Mahasiswa, error) {
	row,err := r.db.Exec("INSERT INTO mahasiswa (id,name,nim,address) VALUES (?,?,?,?)",m.ID,m.Name,m.Nim,m.Address)
	if err != nil {
		return model.Mahasiswa{},err
	}
	id,err := row.LastInsertId()
	if err != nil {
		return model.Mahasiswa{},err
	}
	m.ID = int(id)

	return m,nil
}

func (r *mahasiswaRepo) Update(m model.Mahasiswa) (model.Mahasiswa, error) {
	_,err := r.db.Exec("UPDATE mahasiswa SET name=?,nim=?,address=? WHERE id = ?",m.Name,m.Nim,m.Address,m.ID)
	if err != nil {
		return  model.Mahasiswa{},err
	}
	return m,nil
}

func (r *mahasiswaRepo) Delete(id int) error {	
	result,err := r.db.Exec("DELETE FROM mahasiswa WHERE id =? ",id)
	if err != nil {
		return  err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("data not found")
	}
	return  nil
}
