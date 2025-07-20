package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"station_meteo_api/internal/form"
	"station_meteo_api/internal/model"
	"station_meteo_api/internal/repository"
	"time"
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

func (userService *UserService) Login(data *form.LoginForm) (*model.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Trouver l'utilisateur
	user, err := userService.userRepository.GetUserByUsername(ctx, data.Username)
	if err != nil {
		return nil, errors.New("l'utilisateur n'existe pas")
	}
	if checkPasswordHash(data.Password, user.Password) {
		updateUser, err := userService.createAuthToken(user)
		if err != nil {
			return nil, err
		}
		return updateUser, nil
	} else {
		return nil, errors.New("le mot de passe ne correspond pas")
	}

}

// createAuthToken créer un token d'authentification avec UUID
func (userService *UserService) createAuthToken(user *model.UserModel) (*model.UserModel, error) {
	id := uuid.New().String()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Vérification que 	uuid n'existe pas
	_, err := userService.userRepository.GetUserByAuthToken(ctx, id)
	if err == nil {
		// Appel récursif, car l'id existe déja
		return userService.createAuthToken(user)
	} else {
		user.AuthToken = id
		updateUser, err := userService.userRepository.UpdateUser(ctx, user)
		if err != nil {
			return nil, errors.New("une erreur empêche la génération de la clé d'API")
		}
		return updateUser, nil
	}
}

// hashPassword Hashe le mot de passe courant
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// checkPasswordHash Vérifie un mot de passe en clair et son hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
