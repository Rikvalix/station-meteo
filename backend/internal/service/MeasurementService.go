package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (service *MeasurementService) GetAllMeasurements() ([]model.MeasureModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data, err := service.measureRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *MeasurementService) CreateMeasurement(data map[string]interface{}) (*model.MeasureModel, error) {
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

func parseMeasurementData(data map[string]interface{}) (*model.MeasureModel, error) {
	temp, ok := data["temperature"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'temperature'")
	}

	humidity, ok := data["humidity"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'humidity'")
	}

	address, ok := data["address"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'address'")
	}

	location, ok := data["location"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'location'")
	}

	stationIDStr, ok := data["station_id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'station_id'")
	}

	stationID, err := bson.ObjectIDFromHex(stationIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format for 'station_id'")
	}

	return &model.MeasureModel{
		Date:        time.Now(),
		Temperature: temp,
		Humidity:    humidity,
		Address:     address,
		Location:    location,
		StationID:   stationID,
	}, nil
}
