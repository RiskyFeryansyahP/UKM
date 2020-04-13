package handler

import (
	"encoding/json"
	"log"

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

// RegisterUKM handler to doing registraion ukm for user
func (u *UKMHandler) RegisterUKM(ctx *fasthttp.RequestCtx) {
	var input model.InputRegisterUKM

	body := ctx.Request.Body()

	_ = json.Unmarshal(body, &input)

	ctx.Response.Header.SetContentType("application/json")
	result, err := u.UKMUsecase.Register(ctx, input)
	if err != nil {
		log.Printf("failed register ukm: %v", err.Message)

		_ = json.NewEncoder(ctx).Encode(err)
	}

	_ = json.NewEncoder(ctx).Encode(result)
}
