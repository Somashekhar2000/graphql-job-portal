package service

import (
	"errors"
	"project-gql/graph/model"
	custommodel "project-gql/models"
	"project-gql/repository"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r repository.Users
	c repository.Company
}

func NewService(r repository.Users, c repository.Company) (*Service, error) {
	if r == nil {
		return nil, errors.New("database connection not provided")
	}

	return &Service{r: r, c: c}, nil

}

func (s *Service) UserSignup(nu model.User) (*model.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured while hashing password")
		return &model.User{}, errors.New("hashing password failed")
	}
	user := custommodel.User{UserName: nu.UserName, Email: nu.Email, PasswordHash: string(hashedPass)}

	cu, err := s.r.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("user creation failed")
		return &model.User{}, errors.New("user creation failed")
	}
	user1 := model.User{UserName: cu.UserName, Email: cu.Email, PasswordHash: cu.PasswordHash}
	return &user1, nil

}
func (s *Service) Userlogin(l model.UserLogin) (*model.User, error) {
	fu, err := s.r.GetUserByEmail(l.Email)
	if err != nil {
		log.Error().Err(err).Msg("user login failed")
		return nil, errors.New("user login failed")
	}
	err = bcrypt.CompareHashAndPassword([]byte(fu.PasswordHash), []byte(l.Password))
	if err != nil {
		log.Error().Err(err).Msg("incoorect user password")
		return nil, errors.New("user login failed due to incorrect password")
	}
	ff := model.User{UserName: fu.UserName, Email: fu.Email}
	return &ff, nil

}
