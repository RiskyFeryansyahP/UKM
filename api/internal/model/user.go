package model

import "github.com/confus1on/UKM/ent"

// InputCreateUser is a input model create user
type InputCreateUser struct {
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	Role        int          `json:"role_id"`
	UserProfile *ent.Profile `json:"userprofile"`
}

// InputLoginUser model input to user login
type InputLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
	Status     bool         `json:"status"`
	Result    *ent.User `json:"result_profile"`
}
