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
	User       *ent.User    `json:"result_user"`
	Profile    *ent.Profile `json:"result_profile"`
	Role       *ent.Role    `json:"result_role"`
}
