package structs

import (
	uuid "github.com/google/uuid"
)

type SportmenSport struct {
	Sport string `json:"sport" validate:"oneof=basketball cycling"`
}

type SportmenWithoutId struct {
	Name             string          `json:"name"`
	LastName         string          `json:"last_name"`
	Age              int             `json:"age"`
	Weight           int             `json:"weight"`
	Height           int             `json:"height"`
	CountryBirth     string          `json:"country_birth"`
	CityBirth        string          `json:"city_birth"`
	CountryResidence string          `json:"country_residence"`
	CityResidence    string          `json:"city_residence"`
	LengthResidence  int             `json:"length_residence"`
	Sport            []SportmenSport `json:"sports" validate:"dive"`
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
