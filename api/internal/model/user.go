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
