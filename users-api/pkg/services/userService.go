package services

import (
	"github.com/ashtishad/bookstore/lib"
	"github.com/ashtishad/bookstore/users-api/internal/dto"
	"github.com/ashtishad/bookstore/users-api/pkg/domain"
)

// UserService is our PRIMARY PORT
type UserService interface {
	GetById(id int64) (*dto.UserResponse, lib.RestErr)
	Create(req dto.UserRequest) (*dto.UserResponse, lib.RestErr)
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
