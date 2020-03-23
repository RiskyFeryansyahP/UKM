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
	StatusCode int            `json:"statuscode"`
	Status     bool           `json:"status"`
	Result     ResultRegister `json:"result"`
}

// InputCreateUser is a input model create user
type InputCreateUser struct {
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	UserProfile *ent.Profile `json:"userprofile"`
}
