package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/user"
	"github.com/confus1on/UKM/internal/utils"

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

	_ = json.Unmarshal(body, &input)


	ctx.Response.Header.SetContentType("application/json")

	result, err := u.UserUsecase.CreateUser(context.Background(), input)
	if err != nil {
		log.Printf("Failed Create User %v", err)
		statuscode := http.StatusBadRequest

		ctx.Response.Header.SetStatusCode(statuscode)

		resErr := utils.WrapErrorJson(err, statuscode)
		_ = json.NewEncoder(ctx).Encode(resErr)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}

// LoginUser handle request when user want login
func (u *UserHandler) LoginUser(ctx *fasthttp.RequestCtx) {
	var input model.InputLoginUser

	body := ctx.Request.Body()

	_ = json.Unmarshal(body, &input)

	ctx.Response.Header.SetContentType("application/json")

	result, err := u.UserUsecase.SigninUser(context.Background(), input)
	if err != nil {
		log.Printf("Failed Login User %v", err)
		statuscode := http.StatusUnauthorized

		ctx.Response.Header.SetStatusCode(statuscode)

		resErr := utils.WrapErrorJson(err, statuscode)
		_ = json.NewEncoder(ctx).Encode(resErr)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}
