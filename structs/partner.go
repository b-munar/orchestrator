package structs

import (
	uuid "github.com/google/uuid"
)

type PartnerWithoutId struct {
	Name             string   `json:"name"`
	LastName         string   `json:"last_name"`
	Age              int      `json:"age"`
	Profesion        string   `json:"profession"`
	License          string   `json:"license"`
	CountryBirth     string   `json:"country_birth"`
	CityBirth        string   `json:"city_birth"`
	CountryResidence string   `json:"country_residence"`
	CityResidence    string   `json:"city_residence"`
	Sport            []string `json:"sports"`
	Companies        []string `json:"companies"`
	TypeService      []string `json:"type_services"`
}

type Partner struct {
	UserId uuid.UUID `json:"user"`
	PartnerWithoutId
}

type PartnerResponse struct {
	Partner PartnerWithoutId `json:"partner"`
}
