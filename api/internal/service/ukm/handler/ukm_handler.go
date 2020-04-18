package handler

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/ukm"

	"github.com/valyala/fasthttp"
)

// UKMHandler hold ukm usecase value
type UKMHandler struct {
	UKMUsecase ukm.UsecaseUKM
}

// NewUKMHandler init ukm handler
func NewUKMHandler(UKMUsecase ukm.UsecaseUKM) *UKMHandler {
	return &UKMHandler{UKMUsecase: UKMUsecase}
}

// GetAllUKM handler fetch all data in database
func (u *UKMHandler) GetAllUKM(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentType("application/json")

	result, err := u.UKMUsecase.GetAll(context.Background())
	if err != nil {
		log.Printf("failed get all data ukm: %v", err)

		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}

// RegisterUKM handler doing registraion ukm for user
func (u *UKMHandler) RegisterUKM(ctx *fasthttp.RequestCtx) {
	var input model.InputRegisterUKM

	param := ctx.UserValue("profileID").(string)

	profileID, _ := strconv.Atoi(param)

	body := ctx.Request.Body()

	_ = json.Unmarshal(body, &input)

	ctx.Response.Header.SetContentType("application/json")

	result, err := u.UKMUsecase.RegisterUKM(context.Background(), profileID, input)
	if err != nil {
		log.Printf("failed register ukm: %v", err.Message)

		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}

// UpdateUKM handler update value ukm in database
func (u *UKMHandler) UpdateUKM(ctx *fasthttp.RequestCtx) {
	var input model.InputUpdateUKM

	param := ctx.UserValue("ukmID").(string)

	id, _ := strconv.Atoi(param)

	body := ctx.Request.Body()

	_ = json.Unmarshal(body, &input)

	ctx.Response.Header.SetContentType("application/json")

	result, err := u.UKMUsecase.UpdateUKM(context.Background(), id, input)
	if err != nil {
		log.Printf("failed update ukm: %v", err)

		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}
