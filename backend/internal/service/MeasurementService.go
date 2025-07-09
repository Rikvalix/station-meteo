package service

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"station_meteo_api/internal/form"
	"station_meteo_api/internal/model"
	"station_meteo_api/internal/repository"
	"time"
)

type MeasurementService struct {
	measureRepository *repository.MeasureRepository
}

func NewMeasurementService(repo *repository.MeasureRepository) *MeasurementService {
	return &MeasurementService{
		measureRepository: repo,
	}
}

// CREATE

func (service *MeasurementService) CreateMeasurement(data *form.MeasurementForm) (*model.MeasureModel, error) {
	// Parse et valide les données
	newMeasurement, err := parseMeasurementData(data)

	if err != nil {
		return nil, err
	}

	// Insertion dans MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	created, err := service.measureRepository.Create(ctx, newMeasurement)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// READ

func (service *MeasurementService) GetAllMeasurements(limit int) ([]model.MeasureModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data, err := service.measureRepository.FindAll(ctx, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *MeasurementService) GetMeasureById(id string) (*model.MeasureModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data, err := service.measureRepository.FindByPublicId(ctx, id)
	if err != nil {

		return nil, err
	}
	return data, nil
}

// UPDATE

// DELETE

// UTILS

func parseMeasurementData(data *form.MeasurementForm) (*model.MeasureModel, error) {
	// Générer nouvel id
	id := uuid.New().String()

	return &model.MeasureModel{
		Date:        time.Now(),
		PublicId:    id,
		Temperature: data.Temperature,
		Humidity:    data.Humidity,
		Address:     data.Address,
		Location:    data.Location,
		StationID:   bson.NewObjectID(), // Temporaire
	}, nil
}
