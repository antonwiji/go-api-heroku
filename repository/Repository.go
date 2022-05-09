package repository

import "gorm.io/gorm"

type Repositort interface {
	Create(person Person) (Person, error)
	FindAll() ([]Person, error)
	FindById(ID int) (Person, error)
	Update(person Person) (Person, error)
	Delate(ID int, person Person) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(person Person) (Person, error) {
	err := r.db.Create(&person).Error
	return person, err
}

func (r *repository) FindAll() ([]Person, error) {
	var person []Person
	err := r.db.Find(&person).Error
	return person, err
}

func (r *repository) FindById(ID int) (Person, error) {
	var person Person
	err := r.db.First(&person, ID).Error

	return person, err
}

func (r *repository) Update(person Person) (Person, error) {
	err := r.db.Save(&person).Error
	return person, err
}

func (r *repository) Delate(ID int, person Person) error {
	err := r.db.Delete(&person, ID).Error
	return err
}
