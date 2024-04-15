package structs

import (
	uuid "github.com/google/uuid"
)

type SportmenSport struct {
	Sport string `json:"sport"`
}

type SportmenWithoutId struct {
	Name               string          `json:"name" validate:"required"`
	LastName           string          `json:"last_name" validate:"required"`
	Age                int             `json:"age" validate:"required"`
	Weight             int             `json:"weight" validate:"required"`
	Height             int             `json:"height" validate:"required"`
	Gender             string          `json:"gender" validate:"required"`
	IdentificationType string          `json:"identification_type" validate:"required"`
	Identification     string          `json:"identification" validate:"required"`
	CountryBirth       string          `json:"country_birth" validate:"required"`
	CityBirth          string          `json:"city_birth" validate:"required"`
	CountryResidence   string          `json:"country_residence" validate:"required"`
	CityResidence      string          `json:"city_residence" validate:"required"`
	LengthResidence    int             `json:"length_residence" validate:"required"`
	Sport              []SportmenSport `json:"sports"`
}

type Sportmen struct {
	UserId uuid.UUID `json:"user"`
	SportmenWithoutId
}

type SportmenDelete struct {
	UserId uuid.UUID `json:"user"`
}

type SportmenResponse struct {
	Sportmen SportmenWithoutId `json:"sportmen"`
}
