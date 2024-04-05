package structs

import (
	uuid "github.com/google/uuid"
)

type PartnerWithoutId struct {
	Name               string   `json:"name" validate:"required"`
	LastName           string   `json:"last_name" validate:"required"`
	Age                int      `json:"age" validate:"required"`
	Profesion          string   `json:"profession" validate:"required"`
	License            string   `json:"license"`
	IdentificationType string   `json:"identification_type" validate:"required"`
	Identification     string   `json:"identification" validate:"required"`
	CountryBirth       string   `json:"country_birth" validate:"required"`
	CityBirth          string   `json:"city_birth" validate:"required"`
	CountryResidence   string   `json:"country_residence" validate:"required"`
	CityResidence      string   `json:"city_residence" validate:"required"`
	Sport              []string `json:"sports" validate:"required"`
	Companies          []string `json:"companies" validate:"required"`
	TypeService        []string `json:"type_services" validate:"required"`
}

type Partner struct {
	UserId uuid.UUID `json:"user"`
	PartnerWithoutId
}

type PartnerResponse struct {
	Partner PartnerWithoutId `json:"partner"`
}
