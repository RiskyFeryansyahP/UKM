package model

import "github.com/confus1on/UKM/ent"

// InputRegisterUKM hold value input for register ukm
type InputRegisterUKM struct {
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ProfileID hold the value of the "profile id" field
	ProfileID int `json:"profileid,omitempty"`
}

// ResponseRegisterUKM hold value for response after register ukm
type ResponseRegisterUKM struct {
	StatusCode int          `json:"statuscode"`
	Status     bool         `json:"status"`
	Result     *ent.Profile `json:"result"`
}
