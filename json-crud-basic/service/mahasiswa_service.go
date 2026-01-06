package service

import (
	"errors"

	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/dto"
	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/model"
	"github.com/ahmadzakyrifin/golang-http-projects/json-crud-basic/repo"
)

type MahasiswaService interface {
	FindAll() ([]dto.Mahasiswa, error)
	Create(m dto.Mahasiswa) ( dto.Mahasiswa,error)
	Update(id int,m dto.Mahasiswa) (dto.Mahasiswa,error)
	Delete(id int) error
}

type mahasiswaService struct {
	repo repo.MahasiswaRepository
}

func NewMahasiswaService(r repo.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{repo: r}
}

func (s *mahasiswaService) FindAll() ([]dto.Mahasiswa, error) {
	data, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	dataMahasiswa := make([]dto.Mahasiswa, 0, len(data))
	for _, m := range data {
		dataMahasiswa = append(dataMahasiswa, dto.Mahasiswa{
			ID:      m.ID,
			Name:    m.Name,
			Nim:     m.Nim,
			Address: m.Address,
		})
	}

	return dataMahasiswa, nil
}

func (s *mahasiswaService) Create(req dto.Mahasiswa) (dto.Mahasiswa,error){
	m := model.Mahasiswa{
		ID: req.ID,
		Name: req.Name,
		Nim: req.Nim,
		Address: req.Address,
	}
	data, err := s.repo.Create(m)
	if err != nil{
		return dto.Mahasiswa{},err
	}

	return dto.Mahasiswa{
		ID: data.ID,
		Name: data.Name,
		Nim: data.Nim,
		Address: data.Address,
	},nil
}

func (s *mahasiswaService) Update(id int,req dto.Mahasiswa) (dto.Mahasiswa,error){
	m := model.Mahasiswa{
		ID: id,
		Name: req.Name,
		Nim: req.Nim,
		Address: req.Address,
	}
	data,err := s.repo.Update(m)
	if err != nil {
		return dto.Mahasiswa{},err
	}
	
	return dto.Mahasiswa{
		ID: data.ID,
		Name: data.Name,
		Nim: data.Nim,
		Address: data.Address,
	},nil
}

func (s *mahasiswaService) Delete (id int) (error) {
	if id <= 0 {
		return errors.New("Invalid Id")
	}
	return s.repo.Delete(id)
}