package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"station_meteo_api/internal/form"
	"station_meteo_api/internal/model"
	"station_meteo_api/internal/repository"
	"time"
)

type MeasurementService struct {
	measureRepository *repository.MeasureRepository
	stationRepository *repository.StationRepository
	context           echo.Context
}

func NewMeasurementService(repo *repository.MeasureRepository, stationRepo *repository.StationRepository) *MeasurementService {
	return &MeasurementService{
		measureRepository: repo,
		stationRepository: stationRepo,
	}
}

// CREATE

func (service *MeasurementService) CreateMeasurement(data *form.MeasurementForm, station *model.StationModel) (*model.MeasureModel, error) {
	// Parse et valide les données
	newMeasurement, err := service.parseMeasurementData(data, station)

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

func (service *MeasurementService) GetAllMeasurements(limit int, stationId string) ([]model.MeasureModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	station, err := service.stationRepository.FindByPublicId(ctx, stationId)

	data, err := service.measureRepository.FindAll(ctx, limit, station.ID)
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

func (service *MeasurementService) GetLatestMesure(id string) (*model.MeasureModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stations, err := service.stationRepository.FindByPublicId(ctx, id)

	data, err := service.measureRepository.FindByStationId(ctx, stations.ID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UPDATE

// DELETE

// UTILS

func (service *MeasurementService) parseMeasurementData(data *form.MeasurementForm, station *model.StationModel) (*model.MeasureModel, error) {
	// Générer nouvel id
	id := uuid.New().String()

	return &model.MeasureModel{
		Date:        time.Now(),
		PublicId:    id,
		Temperature: data.Temperature,
		Humidity:    data.Humidity,
		Address:     station.Address,
		Location:    station.Location,
		StationID:   station.ID, // Id de la station émettrice
	}, nil
}
