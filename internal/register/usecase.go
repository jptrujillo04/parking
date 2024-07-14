package register

import (
	"context"
	"errors"
)

type UseCase interface {
	CreateUser(ctx context.Context, userReq UserRequest) error
	GetUser(ctx context.Context, id string) (UserMotorcycles, error)
	UpdateUser(ctx context.Context, userReq UserRequest) error
	GetAllUsers(ctx context.Context) ([]UserMotorcycles, error)
}

type UserUseCase struct {
	Repository RepositoryUser
}

func NewUserUseCase(repository RepositoryUser) *UserUseCase {
	return &UserUseCase{
		Repository: repository,
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, userReq UserRequest) error {
	if userReq.ID == "" {
		return errors.New("invalid user id")
	}
	userCompleted := mapUserRequestToModelUser(userReq)
	return u.Repository.CreateUser(ctx, &userCompleted)
}

func (u *UserUseCase) GetUser(ctx context.Context, id string) (UserMotorcycles, error) {
	return u.Repository.GetUser(ctx, id)
}

func (u *UserUseCase) UpdateUser(ctx context.Context, userReq UserRequest) error {
	if userReq.ID == "" {
		return errors.New("invalid user id")
	}
	userCompleted := mapUserRequestToModelUser(userReq)
	return u.Repository.UpdateUser(ctx, &userCompleted)
}

func (u *UserUseCase) GetAllUsers(ctx context.Context) ([]UserMotorcycles, error) {
	return u.Repository.GetAllUsers(ctx)
}
