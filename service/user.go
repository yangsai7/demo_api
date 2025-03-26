package service

import (
	"context"

	"github.com/yangsai7/demo_api/api"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) Login(ctx context.Context, request *api.UserLoginRequest) (*api.UserLoginResponse, error) {
	return nil, nil
}
