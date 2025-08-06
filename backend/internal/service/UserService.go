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
	userRepository    *repository.UserRepository
	stationRepository *repository.StationRepository
	context.Context
}

func NewUserService(repo *repository.UserRepository, stationRepo *repository.StationRepository) *UserService {
	return &UserService{
		userRepository:    repo,
		stationRepository: stationRepo,
	}
}

func (userService *UserService) Me(userModel *model.UserModel) (*model.UserModel, error) {
	return userModel, nil
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

// GetAllStations Renvoi toutes les stations
func (userService *UserService) GetAllStations() ([]model.StationModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stations, err := userService.stationRepository.FindAll(ctx)
	if err != nil {
		return nil, errors.New("erreur pendant le renvoi des stations")
	}
	return stations, nil
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

func (userService *UserService) UpdatePassword(user *model.UserModel, oldPassword string, newPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oldHash, _ := hashPassword(oldPassword)
	if oldHash != user.Password {
		return errors.New("l'ancien mot de passe ne correspond pas")
	}
	// Check si le nouveau n'est pas l'ancien mot de passe
	newHash, err := hashPassword(newPassword)
	if err != nil {
		return errors.New("erreur empêche le changement du mot de passe")
	}
	if newHash == oldHash {
		return errors.New("l'ancien et le nouveau mot de passe sont similaires")
	}
	user.Password = newHash
	user.AuthToken = "" // Destruction de la session actuelle
	_, err = userService.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return errors.New("erreur d'insertion en base de données")
	}
	return nil
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
