package model

import "github.com/confus1on/UKM/ent"

// InputRegisterUKM hold value input for register ukm
type InputRegisterUKM struct {
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// InputUpdateUKM hold value input for update ukm
type InputUpdateUKM struct {
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// ResponseRegisterUKM hold value for response after register ukm
type ResponseRegisterUKM struct {
	StatusCode int          `json:"statuscode"`
	Status     bool         `json:"status"`
	Result     *ent.Profile `json:"result"`
}

// ResponseGetAllUKM hold data response adter get all ukm
type ResponseGetAllUKM struct {
	StatusCode int        `json:"statuscode"`
	Status     bool       `json:"status"`
	Result     []*ent.Ukm `json:"result"`
}

// ResponseUpdateUKM hold data response after update ukm
type ResponseUpdateUKM struct {
	StatusCode int      `json:"statuscode"`
	Status     bool     `json:"status"`
	Result     *ent.Ukm `json:"result"`
}
