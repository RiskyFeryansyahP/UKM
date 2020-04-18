package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/ukm/mock"
	"github.com/confus1on/UKM/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"net"
	"testing"
)

func TestUKMHandler_GetAllUKM(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	mockUsecase := mock.NewMockUsecaseUKM(controller)

	t.Run("get all ukm handler", func(t *testing.T) {
		response := &model.ResponseGetAllUKM{
			StatusCode: 200,
			Status:     true,
		}

		mockUsecase.EXPECT().GetAll(ctx).Return(response, nil).Times(1)

		ukm := NewUKMHandler(mockUsecase)

		router := fasthttprouter.New()
		router.GET("/api/v0/ukm", ukm.GetAllUKM)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("GET")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/ukm")

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		_ = json.Unmarshal(resp.Body(), &response)

		require.Equal(t, 200, response.StatusCode)
	})

	t.Run("get all ukm handler failed connection refused", func(t *testing.T) {
		respErr := &utils.Error{
			StatusCode: 200,
			Message:    "Connection Refused",
		}

		mockUsecase.EXPECT().GetAll(ctx).Return(nil, respErr).Times(1)

		ukm := NewUKMHandler(mockUsecase)

		router := fasthttprouter.New()
		router.GET("/api/v0/ukm", ukm.GetAllUKM)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("GET")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/ukm")

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		_ = json.Unmarshal(resp.Body(), &respErr)

		require.Equal(t, 200, respErr.StatusCode)
	})
}

func TestUKMHandler_RegisterUKM(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	mockUsecase := mock.NewMockUsecaseUKM(controller)

	t.Run("register ukm handler success", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Name: "SceN",
		}

		profileID := 1

		response := &model.ResponseRegisterUKM{
			StatusCode: 200,
			Status:     true,
		}

		mockUsecase.EXPECT().RegisterUKM(ctx, profileID, input).Return(response, nil).Times(1)

		ukm := NewUKMHandler(mockUsecase)

		router := fasthttprouter.New()
		router.POST("/api/v0/ukm/register/:profileID", ukm.RegisterUKM)

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/ukm/register/1")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err = serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &response)
		require.NoError(t, err)

		require.Equal(t, 200, response.StatusCode)
	})

	t.Run("register ukm handler success", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Name: "",
		}

		profileID := 1

		respErr := &utils.Error{
			StatusCode: 400,
			Message:    "name can't be empty",
		}

		mockUsecase.EXPECT().RegisterUKM(ctx, profileID, input).Return(nil, respErr).Times(1)

		ukm := NewUKMHandler(mockUsecase)

		router := fasthttprouter.New()
		router.POST("/api/v0/ukm/register/:profileID", ukm.RegisterUKM)

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/ukm/register/1")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err = serve(router.Handler, req, resp)
		require.NoError(t, err)

		var response *utils.Error

		err = json.Unmarshal(resp.Body(), &response)
		require.NoError(t, err)

		require.Equal(t, 400, response.StatusCode)
	})
}

func TestUKMHandler_UpdateUKM(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	mockUsecase := mock.NewMockUsecaseUKM(controller)

	t.Run("update ukm handler should success", func(t *testing.T) {
		input := model.InputUpdateUKM{
			Name: "Kompeni",
		}

		ukmID := 1

		res := &model.ResponseUpdateUKM{
			StatusCode: 200,
			Status:     true,
		}

		mockUsecase.EXPECT().UpdateUKM(ctx, ukmID, input).Return(res, nil).Times(1)

		ukm := NewUKMHandler(mockUsecase)

		router := fasthttprouter.New()
		router.PUT("/api/v0/ukm/:ukmID", ukm.UpdateUKM)

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("PUT")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/ukm/1")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err = serve(router.Handler, req, resp)
		require.NoError(t, err)

		var response *model.ResponseUpdateUKM

		_ = json.Unmarshal(resp.Body(), &response)

		require.Equal(t, 200, response.StatusCode)
	})

	t.Run("update ukm handler failed empty name", func(t *testing.T) {
		input := model.InputUpdateUKM{
			Name: "",
		}

		ukmID := 1

		resErr := &utils.Error{
			StatusCode: 400,
			Message:    "name can't be empty",
		}

		mockUsecase.EXPECT().UpdateUKM(ctx, ukmID, input).Return(nil, resErr).Times(1)

		ukm := NewUKMHandler(mockUsecase)

		router := fasthttprouter.New()
		router.PUT("/api/v0/ukm/:ukmID", ukm.UpdateUKM)

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("PUT")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/ukm/1")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err = serve(router.Handler, req, resp)
		require.NoError(t, err)

		var response *utils.Error

		_ = json.Unmarshal(resp.Body(), &response)

		require.Equal(t, 400, response.StatusCode)
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
