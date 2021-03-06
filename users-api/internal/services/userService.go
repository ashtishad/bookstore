package services

import (
	"github.com/ashtishad/bookstore/lib"
	"github.com/ashtishad/bookstore/users-api/internal/domain"
	"github.com/ashtishad/bookstore/users-api/internal/dto"
)

// UserService is our PRIMARY PORT
type UserService interface {
	GetById(id int64) (*dto.UserResponse, lib.RestErr)
	Create(req dto.UserRequest) (*dto.UserResponse, lib.RestErr)
	Update(req dto.UserUpdateRequest) (*dto.UserResponse, lib.RestErr)
	Delete(id int64) lib.RestErr
	SearchByName(name string) (*[]dto.UserSearchResponse, lib.RestErr)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repo}
}

// GetById returns User by id
func (s DefaultUserService) GetById(id int64) (*dto.UserResponse, lib.RestErr) {
	u, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	resp := u.ToUserRespDTO()

	return &resp, nil
}

func (s DefaultUserService) Create(req dto.UserRequest) (*dto.UserResponse, lib.RestErr) {
	//if err := req.Validate(); err != nil {
	//	return nil, err
	//}

	u := domain.NewUser(req)

	usr, err := s.repo.Save(u)
	if err != nil {
		return nil, err
	}

	resp := usr.ToUserRespDTO()

	return &resp, nil
}

func (s DefaultUserService) Update(req dto.UserUpdateRequest) (*dto.UserResponse, lib.RestErr) {
	u := domain.NewUpdateUser(req)

	usr, err := s.repo.Update(u)
	if err != nil {
		return nil, err
	}

	resp := usr.ToUserRespDTO()

	return &resp, nil
}

func (s DefaultUserService) Delete(id int64) lib.RestErr {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

// SearchByName searches users by name
func (s DefaultUserService) SearchByName(name string) (*[]dto.UserSearchResponse, lib.RestErr) {
	users, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.UserSearchResponse, len(*users))
	for i, user := range *users {
		resp[i] = user.ToUserSearchRespDTO()
	}

	return &resp, nil
}
