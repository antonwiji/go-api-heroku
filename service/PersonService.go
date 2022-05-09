package service

import "rest-api-gorm/repository"

type Service interface {
	Create(person PersonRequest) (repository.Person, error)
	FindAll() ([]repository.Person, error)
	FindById(ID int) (repository.Person, error)
	Update(ID int, person PersonUpdate) (repository.Person, error)
	Delete(ID int, person repository.Person) error
}

type service struct {
	repository repository.Repositort
}

func NewService(repository repository.Repositort) *service {
	return &service{repository}
}

func (s *service) Create(person PersonRequest) (repository.Person, error) {

	NewPerson := repository.Person{
		Nama:   person.Nama,
		Alamat: person.Alamat,
	}

	return s.repository.Create(NewPerson)
}

func (s *service) FindAll() ([]repository.Person, error) {

	return s.repository.FindAll()

}

func (s *service) FindById(ID int) (repository.Person, error) {

	return s.repository.FindById(ID)

}

func (s *service) Update(ID int, person PersonUpdate) (repository.Person, error) {

	Newperson, err := s.repository.FindById(ID)
	if err != nil {
		panic(err)
	}

	Newperson.Nama = person.Nama
	Newperson.Alamat = person.Alamat

	PersonNew, err := s.repository.Update(Newperson)

	return PersonNew, err
}

func (s *service) Delete(ID int, person repository.Person) error {
	return s.repository.Delate(ID, person)
}
