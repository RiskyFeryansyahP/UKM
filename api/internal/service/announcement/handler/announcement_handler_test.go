package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/announcement/mock"
	"github.com/confus1on/UKM/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net"
	"testing"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestAnnouncementHandler_GetAnnouncement(t *testing.T) {
	controller := gomock.NewController(t)
	
	ctx := context.Background()
	
	announcementUCMock := mock.NewMockUsecaseAnnouncement(controller)
	
	t.Run("test get announcement route should success", func(t *testing.T) {
		response := &model.ResultAnnouncement{
			StatusCode: 200,
			Status:     true,
		}

		ukmID := 1

		announcementUCMock.EXPECT().GetSpecificAnnouncement(ctx, ukmID).Return(response, nil).Times(1)

		announcement := NewAnnouncementHandler(announcementUCMock)

		router := fasthttprouter.New()
		router.GET("/api/v0/announcement/ukm/:ukmID", announcement.GetAnnouncement)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("GET")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/announcement/ukm/1")

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &response)
		require.NoError(t, err)

		require.Equal(t, 200, response.StatusCode)
	})

	t.Run("test get announcement route failed request time out", func(t *testing.T) {
		responseError := &utils.Error{
			StatusCode: 400,
			Message:    "Request time out",
		}

		ukmID := 1

		announcementUCMock.EXPECT().GetSpecificAnnouncement(ctx, ukmID).Return(nil, responseError).Times(1)

		announcement := NewAnnouncementHandler(announcementUCMock)

		router := fasthttprouter.New()
		router.GET("/api/v0/announcement/ukm/:ukmID", announcement.GetAnnouncement)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("GET")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/announcement/ukm/1")

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &responseError)
		require.NoError(t, err)

		require.Equal(t, 400, responseError.StatusCode)
	})
}

func TestAnnouncementHandler_PostAnnouncement(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	announcementUCMock := mock.NewMockUsecaseAnnouncement(controller)

	t.Run("test post announcement should success", func(t *testing.T) {
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		ukmID := 1

		response := &model.SingleResultAnnouncement{
			StatusCode: 200,
			Status:     true,
		}

		announcementUCMock.EXPECT().ValidationPostAnnouncement(ctx, ukmID, input).Return(response, nil).Times(1)

		announcement := NewAnnouncementHandler(announcementUCMock)

		router := fasthttprouter.New()
		router.POST("/api/v0/announcement/ukm/:ukmID", announcement.PostAnnouncement)

		body, _ := json.Marshal(input)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/announcement/ukm/1")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &response)
		require.NoError(t, err)
		require.Equal(t, 200, response.StatusCode)
	})

	t.Run("test post announcement failed empty title", func(t *testing.T) {
		input := model.InputAnnouncement{
			Title:       "",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		ukmID := 1

		responseError := &utils.Error{
			StatusCode: 400,
			Message:    "title can't be null or empty",
		}

		announcementUCMock.EXPECT().ValidationPostAnnouncement(ctx, ukmID, input).Return(nil, responseError).Times(1)

		announcement := NewAnnouncementHandler(announcementUCMock)

		router := fasthttprouter.New()
		router.POST("/api/v0/announcement/ukm/:ukmID", announcement.PostAnnouncement)

		body, _ := json.Marshal(input)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/announcement/ukm/1")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &responseError)
		require.NoError(t, err)
		require.Equal(t, 400, responseError.StatusCode)
	})
}

func TestAnnouncementHandler_UpdateAnnouncement(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	announcementUCMock := mock.NewMockUsecaseAnnouncement(controller)

	t.Run("test update announcement should success", func(t *testing.T) {
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 2",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		announcementID := 1

		response := &model.SingleResultAnnouncement{
			StatusCode: 200,
			Status:     true,
		}

		announcementUCMock.EXPECT().UpdateOneAnnouncement(ctx, announcementID, input).Return(response, nil).Times(1)

		announcement := NewAnnouncementHandler(announcementUCMock)

		router := fasthttprouter.New()
		router.PUT("/api/v0/announcement/:announcementID", announcement.UpdateAnnouncement)

		body, _ := json.Marshal(input)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("PUT")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/announcement/1")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &response)
		require.NoError(t, err)
		require.Equal(t, 200, response.StatusCode)
	})

	t.Run("test update announcement failed zero announcement id", func(t *testing.T) {
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 2",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		announcementID := 0

		responseError := &utils.Error{
			StatusCode: 400,
			Message:    "announcement id can't be null or zero",
		}

		announcementUCMock.EXPECT().UpdateOneAnnouncement(ctx, announcementID, input).Return(nil, responseError).Times(1)

		announcement := NewAnnouncementHandler(announcementUCMock)

		router := fasthttprouter.New()
		router.PUT("/api/v0/announcement/:announcementID", announcement.UpdateAnnouncement)

		body, _ := json.Marshal(input)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("PUT")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/announcement/0")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &responseError)
		require.NoError(t, err)
		require.Equal(t, 400, responseError.StatusCode)
	})
}

func TestAnnouncementHandler_DeleteAnnouncement(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	announcementUCMock := mock.NewMockUsecaseAnnouncement(controller)

	t.Run("test delete announcement should success", func(t *testing.T) {
		announcementID := 1

		response := &model.SingleResultAnnouncement{
			StatusCode: 200,
			Status:     true,
		}

		announcementUCMock.EXPECT().DeleteOneAnnouncement(ctx, announcementID).Return(response, nil).Times(1)

		announcement := NewAnnouncementHandler(announcementUCMock)

		router := fasthttprouter.New()
		router.DELETE("/api/v0/announcement/:announcementID", announcement.DeleteAnnouncement)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("DELETE")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/announcement/1")

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &response)
		require.NoError(t, err)
		require.Equal(t, 200, response.StatusCode)
	})

	t.Run("test delete announcement failed not found announcement id", func(t *testing.T) {
		announcementID := 1002

		responseError := &utils.Error{
			StatusCode: 400,
			Message:    "Not found announcement id",
		}

		announcementUCMock.EXPECT().DeleteOneAnnouncement(ctx, announcementID).Return(nil, responseError).Times(1)

		announcement := NewAnnouncementHandler(announcementUCMock)

		router := fasthttprouter.New()
		router.DELETE("/api/v0/announcement/:announcementID", announcement.DeleteAnnouncement)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("DELETE")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/api/v0/announcement/1002")

		resp := fasthttp.AcquireResponse()

		err := serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &responseError)
		require.NoError(t, err)
		require.Equal(t, 400, responseError.StatusCode)
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
