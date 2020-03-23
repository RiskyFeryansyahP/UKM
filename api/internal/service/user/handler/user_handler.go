package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/user"

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
		log.Fatalf("Failed unmarshan body %v", err)
		ctx.Response.Header.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBodyRaw([]byte(`{"Message", "` + err.Error() + `"}`))
		return
	}

	result, err := u.UserUsecase.CreateUser(ctx, input)
	if err != nil {
		log.Printf("Failed Create User %v", err)
		ctx.Response.Header.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.SetBodyRaw([]byte(`{"Message", "` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(ctx).Encode(result)
}
