package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"project-gql/graph/model"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	PasswordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := model.User{UserName: input.UserName, Email: input.Email, PasswordHash: string(PasswordHash)}
	return r.S.UserSignup(user)
}

// CreateCompany is the resolver for the createCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input model.NewCompany) (*model.Company, error) {
	return r.S.CompanyCreate(input)
}

// CreateJobPosting is the resolver for the createJobPosting field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*model.Job, error) {
	return r.S.JobCreate(input)
}

// CheckUser is the resolver for the checkUser field.
func (r *queryResolver) CheckUser(ctx context.Context, input model.UserLogin) (*model.User, error) {
	return r.S.Userlogin(input)
}

// GetCompanies is the resolver for the getCompanies field.
func (r *queryResolver) GetCompanies(ctx context.Context) ([]*model.Company, error) {
	return r.S.GetAllCompanies()
}

// GetCompany is the resolver for the getCompany field.
func (r *queryResolver) GetCompany(ctx context.Context, input string) (*model.Company, error) {
	return r.S.GetCompany(input)
}

// GetJobs is the resolver for the getJobs field.
func (r *queryResolver) GetJobs(ctx context.Context, input string) ([]*model.Job, error) {
	return r.S.GetJobs(input)
}

// GetAllJobs is the resolver for the getAllJobs field.
func (r *queryResolver) GetAllJobs(ctx context.Context) ([]*model.Job, error) {
	return r.S.GetAllJobs()
}

// GetJobByID is the resolver for the getJobById field.
func (r *queryResolver) GetJobByID(ctx context.Context, input string) (*model.Job, error) {
	panic(fmt.Errorf("not implemented: GetJobByID - getJobById"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: FindUserByEmail - findUserByEmail"))
}