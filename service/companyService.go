package service

import (
	"errors"
	"project-gql/graph/model"
	custommodel "project-gql/models"
	"strconv"

	"github.com/rs/zerolog/log"
)

func (s *Service) CompanyCreate(mnc model.NewCompany) (*model.Company, error) {
	company := custommodel.Company{CompanyName: mnc.CompanyName, Adress: mnc.Adress, Domain: mnc.Domain}
	cc, err := s.c.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("user creation failed")
		return &model.Company{}, errors.New("user creation failed")
	}
	u := model.Company{CompanyName: cc.CompanyName, Adress: cc.Adress, Domain: cc.Domain}
	return &u, nil
}

func (s *Service) GetAllCompanies() ([]*model.Company, error) {

	AllCompanies, err := s.c.GetAllCompany()
	if err != nil {
		return nil, err
	}
	var Companies []*model.Company

	for _, value := range AllCompanies {
		Company := model.Company{CompanyName: value.CompanyName, Adress: value.Adress, Domain: value.Domain}
		Companies = append(Companies, &Company)
	}
	return Companies, nil

}

func (s *Service) GetCompany(id string) (*model.Company, error) {
	intValue, _ := strconv.Atoi(id)
	Company, err := s.c.GetCompanyById(intValue)
	if err != nil {
		log.Error().Err(err).Msg("error while parsing id")
		return nil, err
	}
	CustomCompany := model.Company{CompanyName: Company.CompanyName, Adress: Company.Adress, Domain: Company.Domain}
	return &CustomCompany, nil

}

func (s *Service) JobCreate(nj model.NewJob) (*model.Job, error) {
	id, err := strconv.ParseUint(nj.ID, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("error while parsing id")
		return nil, err
	}

	job := custommodel.Job{JobTitle: nj.JobTitle, JobSalary: nj.JobSalary, Uid: id}
	cu, err := s.c.CreateJob(job)
	if err != nil {
		log.Error().Err(err).Msg("job creation failed")
		return nil, errors.New("job creation failed")
	}
	str := strconv.FormatUint(cu.Uid, 10)
	company, err := s.GetCompany(str)
	if err != nil {
		log.Error().Err(err).Msg("error while retriving company")
		return nil, errors.New("error while retriving company")
	}
	cugr := model.Job{JobTitle: cu.JobTitle, JobSalary: cu.JobSalary, Company: company}

	return &cugr, nil
}

func (s *Service) GetJobs(id string) ([]*model.Job, error) {
	intValue, _ := strconv.Atoi(id)
	AllCompanies, err := s.c.GetJobById(intValue)
	if err != nil {
		return nil, errors.New("job retreval failed")
	}
	var Jobs []*model.Job

	for _, value := range AllCompanies {
		Job := model.Job{JobTitle: value.JobTitle, JobSalary: value.JobSalary}
		Jobs = append(Jobs, &Job)
	}
	return Jobs, nil
}

func (s *Service) GetAllJobs() ([]*model.Job, error) {

	AllJobs, err := s.c.GetAllJobs()
	if err != nil {
		return nil, err
	}
	var Jobs []*model.Job
	for _, value := range AllJobs {
		Job := model.Job{JobTitle: value.JobTitle, JobSalary: value.JobSalary}
		Jobs = append(Jobs, &Job)
	}
	return Jobs, nil

}
