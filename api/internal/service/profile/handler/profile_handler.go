package handler

import (
	"encoding/json"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile"
	"github.com/confus1on/UKM/internal/utils"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"strconv"
)

type ProfileHandler struct {
	ProfileUsecase profile.UsecaseProfile
}

func NewProfileHandler(profileUsecase profile.UsecaseProfile) *ProfileHandler {
	return &ProfileHandler{ProfileUsecase: profileUsecase}
}

func (p *ProfileHandler) UpdateProfile(ctx *fasthttp.RequestCtx)  {
	var input model.InputUpdateProfile

	body := ctx.Request.Body()

	_ = json.Unmarshal(body, &input)

	param := ctx.UserValue("id").(string)
	id, _ := strconv.Atoi(param)

	resp, err := p.ProfileUsecase.UpdateProfile(ctx, id, input)
	if err != nil {
		log.Println("failed update profile")

		statuscode := http.StatusBadRequest

		ctx.Response.Header.SetStatusCode(statuscode)

		respErr := utils.WrapErrorJson(err, statuscode)

		_ = json.NewEncoder(ctx).Encode(respErr)
	}

	_ = json.NewEncoder(ctx).Encode(resp)
}