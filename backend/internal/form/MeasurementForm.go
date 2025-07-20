package form

type (
	MeasurementForm struct {
		Temperature float64 `json:"temperature" validate:"required"` // Temperature
		Humidity    float64 `json:"humidity" validate:"required"`    // Humidité
	}
)
