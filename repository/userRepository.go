package repository

import (
	"errors"
	custommodel "project-gql/models"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (*Repo, error) {
	if db == nil {
		return nil, errors.New("db connection not given")
	}

	return &Repo{db: db}, nil

}

type Users interface {
	CreateUser(cu custommodel.User) (custommodel.User, error)
	GetUserByEmail(s string) (*custommodel.User, error)
}

func (r *Repo) CreateUser(cu custommodel.User) (custommodel.User, error) {

	err := r.db.Create(&cu).Error
	if err != nil {
		return custommodel.User{}, err
	}
	return cu, nil
}
func (r *Repo) GetUserByEmail(s string) (*custommodel.User, error) {
	var cu custommodel.User
	tx := r.db.Where("email=?", s).First(&cu)
	if tx.Error != nil {
		return nil, nil
	}
	return &cu, nil

}
