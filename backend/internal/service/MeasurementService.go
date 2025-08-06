package service

import (
	"context"
	"errors"
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
	if err != nil {
		return nil, errors.New("failed to find station")
	}
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

func (service *MeasurementService) GetMeasuresByDate(date string) ([]model.MeasureModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Vérifier si la date est valide
	if !service.checkDate(date) {
		return nil, errors.New("date format is not equal to: YYYY-MM-DD")
	}
	dateConvert, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("error to parse date")
	}
	results, err := service.measureRepository.FindBySpecificDay(ctx, dateConvert)
	if err != nil {
		return nil, errors.New("error to find measurements by date")
	}
	if (results == nil) || (len(results) == 0) {
		return nil, errors.New("error to find measurements by date")
	}
	return results, nil

}

// UPDATE

// DELETE

// UTILS

// Vérifier si la date est format YYYY-MM-DD
func (service *MeasurementService) checkDate(date string) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, date)
	return err == nil
}

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
