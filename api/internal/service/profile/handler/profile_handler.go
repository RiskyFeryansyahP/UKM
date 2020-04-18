package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile"
	"github.com/valyala/fasthttp"
)

// ProfileHandler model of handler profile
type ProfileHandler struct {
	ProfileUsecase profile.UsecaseProfile
}

// NewProfileHandler create a new profile handler to REST
func NewProfileHandler(profileUsecase profile.UsecaseProfile) *ProfileHandler {
	return &ProfileHandler{ProfileUsecase: profileUsecase}
}

// GetDetailProfile get detail of a profile handler
func (p *ProfileHandler) GetDetailProfile(ctx *fasthttp.RequestCtx) {
	email := ctx.UserValue("email").(string)


	ctx.Response.Header.SetContentType("application/json")

	resp, err := p.ProfileUsecase.GetDetailProfile(context.Background(), email)
	if err != nil {
		log.Printf("failed get detail profile %v", err)

		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(resp)
}

// UpdateProfile update a profile handler
func (p *ProfileHandler) UpdateProfile(ctx *fasthttp.RequestCtx) {
	var input model.InputUpdateProfile

	body := ctx.Request.Body()

	_ = json.Unmarshal(body, &input)

	email := ctx.UserValue("email").(string)


	ctx.Response.Header.SetContentType("application/json")

	resp, err := p.ProfileUsecase.UpdateProfile(context.Background(), email, input)
	if err != nil {
		log.Println("failed update profile")

		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(resp)
}
