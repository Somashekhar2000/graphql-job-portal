package repository

import custommodel "project-gql/models"

type Company interface {
	CreateCompany(custommodel.Company) (custommodel.Company, error)
	GetAllCompany() ([]custommodel.Company, error)
	GetCompanyById(id int) (custommodel.Company, error)
	CreateJob(j custommodel.Job) (custommodel.Job, error)
	GetJobById(id int) ([]custommodel.Job, error)
	GetAllJobs() ([]custommodel.Job, error)
}

func (r *Repo) CreateCompany(cc custommodel.Company) (custommodel.Company, error) {
	err := r.db.Create(&cc).Error
	if err != nil {
		return custommodel.Company{}, err
	}
	return cc, nil
}

func (r *Repo) GetAllCompany() ([]custommodel.Company, error) {
	var cc []custommodel.Company
	err := r.db.Find(&cc).Error
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (r *Repo) GetCompanyById(id int) (custommodel.Company, error) {
	var cc custommodel.Company

	tx := r.db.Where("id = ?", id)
	err := tx.Find(&cc).Error
	if err != nil {
		return custommodel.Company{}, err
	}
	return cc, nil

}

func (r *Repo) CreateJob(j custommodel.Job) (custommodel.Job, error) {
	err := r.db.Create(&j).Error
	if err != nil {
		return custommodel.Job{}, err
	}
	return j, nil
}

func (r *Repo) GetJobById(id int) ([]custommodel.Job, error) {
	var cj []custommodel.Job

	tx := r.db.Where("uid = ?", id)
	err := tx.Find(&cj).Error
	if err != nil {
		return nil, err
	}
	return cj, nil

}

func (r *Repo) GetAllJobs() ([]custommodel.Job, error) {
	var cj []custommodel.Job
	err := r.db.Find(&cj).Error
	if err != nil {
		return nil, err
	}

	return cj, nil
}
