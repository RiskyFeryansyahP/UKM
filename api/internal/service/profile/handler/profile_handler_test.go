package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"testing"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile/mock"
	"github.com/confus1on/UKM/internal/utils"

	"github.com/buaazp/fasthttprouter"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestProfileHandler_GetDetailProfile(t *testing.T) {
	controller := gomock.NewController(t)

	t.Run("get detail profile handler should success", func(t *testing.T) {
		email := "171111040@mhs.stiki.ac.id"

		profile := &ent.Profile{
			ID:             1,
			FirstName:      "Galih",
			LastName:       "Satriawan",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		response := &model.ResponseGetProfile{
			StatusCode: 200,
			Status:     true,
			Result:     profile,
		}

		uc := mock.NewMockUsecaseProfile(controller)
		uc.EXPECT().GetDetailProfile(context.Background(), email).Return(response, nil).Times(2)

		p := NewProfileHandler(uc)

		router := fasthttprouter.New()
		router.GET("/:email", p.GetDetailProfile)

		req := fasthttp.AcquireRequest()
		req.SetRequestURI("/" + email)
		req.Header.SetMethod("GET")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")

		resp := fasthttp.AcquireResponse()
		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &response)
		require.NoError(t, err)

		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, true, response.Status)
	})

	t.Run("get detail profile handler should failed wrong email", func(t *testing.T) {
		email := "171111020@mhs.stiki.ac.id"

		errs := &utils.Error{
			StatusCode: 200,
			Message:    "failed get detail profile",
		}

		uc := mock.NewMockUsecaseProfile(controller)
		uc.EXPECT().GetDetailProfile(context.Background(), email).Return(nil, errs).Times(1)

		p := NewProfileHandler(uc)

		router := fasthttprouter.New()
		router.GET("/:email", p.GetDetailProfile)

		req := fasthttp.AcquireRequest()
		req.SetRequestURI("/" + email)
		req.Header.SetMethod("GET")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")

		resp := fasthttp.AcquireResponse()
		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		resErr := utils.Error{}

		err = json.Unmarshal(resp.Body(), &resErr)

		require.NoError(t, err)
		require.Equal(t, 200, resErr.StatusCode)
	})
}

func TestProfileHandler_UpdateProfile(t *testing.T) {
	controller := gomock.NewController(t)

	t.Run("update profile handler should be success", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			YearGeneration: "2016",
			Phone:          "089691952593",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		profile := &ent.Profile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			YearGeneration: "2016",
			Phone:          "089691952593",
			Status:         true,
			ImgProfile:     "",
		}

		response := &model.ResponseUpdateProfile{
			StatusCode: 200,
			Status:     true,
			Result:     profile,
		}

		uc := mock.NewMockUsecaseProfile(controller)
		uc.EXPECT().UpdateProfile(context.Background(), email, input).Return(response, nil).Times(1)

		p := NewProfileHandler(uc)

		router := fasthttprouter.New()
		router.PUT("/api/v0/profile/update/:email", p.UpdateProfile)

		b, err := json.Marshal(input)
		require.NoError(t, err)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("PUT")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/profile/update/" + email)
		req.SetBody(b)

		resp := fasthttp.AcquireResponse()
		err = serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &response)
		require.NoError(t, err)

		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, true, response.Status)
	})

	t.Run("update profile handler failed data input", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "",
			LastName:       "",
			YearGeneration: "2016",
			Phone:          "089691952593",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		errs := &utils.Error{
			StatusCode: 400,
			Message:    "firstname and lastname can't be empty",
		}

		uc := mock.NewMockUsecaseProfile(controller)
		uc.EXPECT().UpdateProfile(context.Background(), email, input).Return(nil, errs).Times(1)

		p := NewProfileHandler(uc)

		router := fasthttprouter.New()
		router.PUT("/api/v0/profile/update/:email", p.UpdateProfile)

		b, err := json.Marshal(input)
		require.NoError(t, err)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("PUT")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/profile/update/" + email)
		req.SetBody(b)

		resp := fasthttp.AcquireResponse()
		err = serve(router.Handler, req, resp)
		require.NoError(t, err)

		respErr := new(utils.Error)

		err = json.Unmarshal(resp.Body(), &respErr)
		require.NoError(t, err)

		require.Equal(t, 400, respErr.StatusCode)
	})
}

// serve serves http request using provided fasthttp handler
func serve(handler fasthttp.RequestHandler, req *fasthttp.Request, res *fasthttp.Response) error {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	go func() {
		err := fasthttp.Serve(ln, handler)
		if err != nil {
			panic(fmt.Errorf("failed to serve: %v", err))
		}
	}()

	client := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	return client.Do(req, res)
}
