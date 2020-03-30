package model

import (
	"github.com/confus1on/UKM/ent"
)

// ResultRegister is model to show User after register
type ResultRegister struct {
	User *ent.User `json:"user"`
}

// ResponseRegister is model Response after successfull register
type ResponseRegister struct {
	StatusCode int       `json:"statuscode"`
	Status     bool      `json:"status"`
	Result     *ent.User `json:"result"`
}

// ResponseLogin model of response after login
type ResponseLogin struct {
	StatusCode int          `json:"statuscode"`
	Status     bool         `json:""`
	Profile    *ent.Profile `json:"profile"`
	Role       *ent.Role    `json:"role"`
}
