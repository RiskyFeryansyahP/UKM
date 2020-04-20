package model

import (
	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/ent/profile"
)

// InputUpdateProfile is the model update entity for the Profile schema.
type InputUpdateProfile struct {
	// FirstName holds the value of the "firstName" field.
	FirstName string `json:"firstName,omitempty"`
	// LastName holds the value of the "lastName" field.
	LastName string `json:"lastName,omitempty"`
	// Jk holds the value of the "jk" field.
	Jk profile.Jk `json:"jk,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate string `json:"birth_date,omitempty"`
	// YearGeneration holds the value of the "year_generation" field.
	YearGeneration string `json:"year_generation,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// ImgProfile holds the value of the "img_profile" field.
	ImgProfile string `json:"img_profile,omitempty"`
}

// ResponseUpdateProfile hold value response after update profile
type ResponseUpdateProfile struct {
	StatusCode int          `json:"statuscode"`
	Status     bool         `json:"status"`
	Result     *ent.Profile `json:"result"`
}

// ResponseGetProfile hold value after get profile
type ResponseGetProfile struct {
	StatusCode int          `json:"statuscode"`
	Status     bool         `json:"status"`
	Result     *ent.Profile `json:"result"`
}
