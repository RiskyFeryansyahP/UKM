package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile"
	"github.com/confus1on/UKM/internal/utils"

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
	ctx.SetUserValue("email", "")
	email := ctx.UserValue("email").(string)

	resp, err := p.ProfileUsecase.GetProfile(ctx, email)
	if err != nil {
		log.Println("failed get detail profile")

		statuscode := http.StatusBadRequest

		respErr := utils.WrapErrorJson(err, statuscode)

		_ = json.NewEncoder(ctx).Encode(respErr)
	}

	_ = json.NewEncoder(ctx).Encode(resp)
}

// UpdateProfile update a profile handler
func (p *ProfileHandler) UpdateProfile(ctx *fasthttp.RequestCtx) {
	var input model.InputUpdateProfile

	body := ctx.Request.Body()

	_ = json.Unmarshal(body, &input)

	email := ctx.UserValue("email").(string)

	resp, err := p.ProfileUsecase.UpdateProfile(ctx, email, input)
	if err != nil {
		log.Println("failed update profile")

		statuscode := http.StatusBadRequest

		ctx.Response.Header.SetStatusCode(statuscode)

		respErr := utils.WrapErrorJson(err, statuscode)

		_ = json.NewEncoder(ctx).Encode(respErr)
	}

	_ = json.NewEncoder(ctx).Encode(resp)
}
