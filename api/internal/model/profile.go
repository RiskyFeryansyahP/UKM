package model

import "github.com/confus1on/UKM/ent"

// InputUpdateProfile is the model update entity for the Profile schema.
type InputUpdateProfile struct {
	// FirstName holds the value of the "firstName" field.
	FirstName string `json:"firstName,omitempty"`
	// LastName holds the value of the "lastName" field.
	LastName string `json:"lastName,omitempty"`
	// YearGeneration holds the value of the "year_generation" field.
	YearGeneration string `json:"year_generation,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// ImgProfile holds the value of the "img_profile" field.
	ImgProfile string `json:"img_profile,omitempty"`
}

type ResponseUpdateProfile struct {
	StatusCode int          `json:"statuscode"`
	Status     bool         `json:"status"`
	Result     *ent.Profile `json:"result"`
}

type ResponseGetProfile struct {
	StatusCode int          `json:"statuscode"`
	Status     bool         `json:"status"`
	Result     *ent.Profile `json:"result"`
}
