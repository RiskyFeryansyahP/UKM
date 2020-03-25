package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RiskyFeryansyahP/UKM/internal/model"
	"github.com/RiskyFeryansyahP/UKM/internal/service/user"
	"github.com/RiskyFeryansyahP/UKM/internal/utils"

	"github.com/valyala/fasthttp"
)

// UserHandler struct injected from user usecase depedency
type UserHandler struct {
	UserUsecase user.UsecaseUser
}

// NewUserHandler is constructor to initialize UserHandler and return UserHandler
func NewUserHandler(userUsecase user.UsecaseUser) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

// RegisterUser is call create user in user usecase
func (u *UserHandler) RegisterUser(ctx *fasthttp.RequestCtx) {
	var input model.InputCreateUser

	body := ctx.Request.Body()

	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Printf("Failed unmarshan body %v", err)

		statuscode := http.StatusInternalServerError

		ctx.Response.Header.SetStatusCode(statuscode)

		resErr := utils.WrapErrorJson(err, statuscode)
		_ = json.NewEncoder(ctx).Encode(resErr)
		return
	}

	result, err := u.UserUsecase.CreateUser(ctx, input)
	if err != nil {
		log.Printf("Failed Create User %v", err)
		statuscode := http.StatusBadRequest

		ctx.Response.Header.SetStatusCode(statuscode)

		resErr := utils.WrapErrorJson(err, statuscode)
		_ = json.NewEncoder(ctx).Encode(resErr)
		return
	}

	json.NewEncoder(ctx).Encode(result)
}

// LoginUser handle request when user want login
func (u *UserHandler) LoginUser(ctx *fasthttp.RequestCtx) {
	var input model.InputLoginUser

	body := ctx.Request.Body()

	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Printf("Failed unmarshal body %v", err)

		statuscode := http.StatusInternalServerError

		ctx.Response.Header.SetStatusCode(statuscode)

		resErr := utils.WrapErrorJson(err, statuscode)
		_ = json.NewEncoder(ctx).Encode(resErr)
		return
	}

	result, err := u.UserUsecase.SigninUser(ctx, input)
	if err != nil {
		log.Printf("Failed Login User %v", err)
		statuscode := http.StatusBadRequest

		ctx.Response.Header.SetStatusCode(statuscode)

		resErr := utils.WrapErrorJson(err, statuscode)
		_ = json.NewEncoder(ctx).Encode(resErr)
		return
	}

	json.NewEncoder(ctx).Encode(result)
}
