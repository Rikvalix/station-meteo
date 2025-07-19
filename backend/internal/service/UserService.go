package service

import (
	"context"
	"station_meteo_api/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
	context.Context
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}
