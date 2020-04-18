package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/user/mock"
	"github.com/confus1on/UKM/internal/utils"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestUserHandler_RegisterUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("test http create user with valid data", func(t *testing.T) {
		profile := &ent.Profile{
			ID:             1,
			FirstName:      "Risky F",
			LastName:       "Pribadi",
			YearGeneration: "2017",
			Phone:          "083834121712",
			Status:         true,
			ImgProfile:     "",
		}

		input := model.InputCreateUser{
			Email:       "171111040@mhs.stiki.ac.id",
			Password:    "risky",
			UserProfile: profile,
		}

		resp := &model.ResponseRegister{
			StatusCode: 200,
			Status:     true,
			Result: &ent.User{
				Email:    input.Email,
				Password: input.Password,
			},
		}

		u := mock.NewMockUsecaseUser(controller)
		u.EXPECT().CreateUser(context.Background(), input).Return(resp, nil).Times(1)

		user := NewUserHandler(u)

		b, _ := json.Marshal(input)

		r, err := http.NewRequest("POST", "http://localhost/", bytes.NewBuffer(b))
		require.NoError(t, err)

		res, err := serve(user.RegisterUser, r)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)

		err = json.Unmarshal(body, &resp)
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
		require.Equal(t, true, resp.Status)
	})

	t.Run("test http create user with invalid data", func(t *testing.T) {
		profile := &ent.Profile{
			ID:             1,
			FirstName:      "Risky F",
			LastName:       "Pribadi",
			YearGeneration: "2017",
			Phone:          "083834121712",
			Status:         true,
			ImgProfile:     "",
		}

		input := model.InputCreateUser{
			Email:       "d",
			Password:    "risky",
			UserProfile: profile,
		}

		resp := utils.Error{}

		err := errors.New("error invalid data")

		u := mock.NewMockUsecaseUser(controller)
		u.EXPECT().CreateUser(context.Background(), input).Return(nil, err).Times(1)

		user := NewUserHandler(u)

		b, _ := json.Marshal(input)

		r, err := http.NewRequest("POST", "http://localhost/", bytes.NewBuffer(b))
		require.NoError(t, err)

		res, err := serve(user.RegisterUser, r)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)

		err = json.Unmarshal(body, &resp)
		require.NoError(t, err)
		require.Equal(t, 400, resp.StatusCode)
	})
}

func TestUserHandler_LoginUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("test http login with valid data", func(t *testing.T) {
		profile := &ent.Profile{
			ID:             1,
			FirstName:      "Risky F",
			LastName:       "Pribadi",
			YearGeneration: "2017",
			Phone:          "083834121712",
			Status:         true,
			ImgProfile:     "",
		}

		role := &ent.Role{
			ID:    1,
			Value: "DEVELOPER",
		}

		u := &ent.User{
			ID:       1,
			Email:    "171111040@mhs.stiki.ac.id",
			Password: "risky",
			Edges:    ent.UserEdges{
				Profile: profile,
				Role:    role,
			},
		}

		input := model.InputLoginUser{
			Email:    "171111040@mhs.stiki.ac.id",
			Password: "risky",
		}

		resp := &model.ResponseLogin{
			StatusCode: 200,
			Status:     true,
			Result: u,
		}

		userUC := mock.NewMockUsecaseUser(controller)
		userUC.EXPECT().SigninUser(context.Background(), input).Return(resp, nil).Times(1)

		user := NewUserHandler(userUC)

		b, _ := json.Marshal(input)

		r, err := http.NewRequest("POST", "http://localhost/", bytes.NewBuffer(b))
		require.NoError(t, err)

		res, err := serve(user.LoginUser, r)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)

		err = json.Unmarshal(body, &resp)
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
		require.Equal(t, true, resp.Status)
	})

	t.Run("test http login with wrong data", func(t *testing.T) {
		input := model.InputLoginUser{
			Email:    "171111040@mhs.stiki.ac.id",
			Password: "risky",
		}

		resp := utils.Error{}

		err := errors.New("wrong username password")

		u := mock.NewMockUsecaseUser(controller)
		u.EXPECT().SigninUser(context.Background(), input).Return(nil, err).Times(1)

		user := NewUserHandler(u)

		b, _ := json.Marshal(input)

		r, err := http.NewRequest("POST", "http://localhost/", bytes.NewBuffer(b))
		require.NoError(t, err)

		res, err := serve(user.LoginUser, r)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)

		err = json.Unmarshal(body, &resp)
		require.NoError(t, err)
		require.Equal(t, 401, resp.StatusCode)
	})
}

// serve serves http request using provided fasthttp handler
func serve(handler fasthttp.RequestHandler, req *http.Request) (*http.Response, error) {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	go func() {
		err := fasthttp.Serve(ln, handler)
		if err != nil {
			panic(fmt.Errorf("failed to serve: %v", err))
		}
	}()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return ln.Dial()
			},
		},
	}

	return client.Do(req)
}
