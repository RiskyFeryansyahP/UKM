package usecase

import (
	"context"
	"errors"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/announcement"
	"github.com/confus1on/UKM/internal/utils"

	"github.com/valyala/fasthttp"
)

// AnnouncementUsecase hold announcement repository value
type AnnouncementUsecase struct {
	AnnouncementRepo announcement.RepositoryAnnouncement
}

// NewAnnouncementUsecase initiate announcement usecase to inject announcement repository
func NewAnnouncementUsecase(announcementRepo announcement.RepositoryAnnouncement) announcement.UsecaseAnnouncement {
	return &AnnouncementUsecase{AnnouncementRepo: announcementRepo}
}

// GetSpecificAnnouncement get specific of announcement based on ukm you want
// return new result announcement and error from utils
func (a *AnnouncementUsecase) GetSpecificAnnouncement(ctx context.Context, ukmID int) (*model.ResultAnnouncement, *utils.Error) {
	if ukmID == 0 || &ukmID == nil {
		responseError := errors.New("ukm id can't be null or zero")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	announcements, err := a.AnnouncementRepo.GetUkmAnnouncement(ctx, ukmID)
	if err != nil {
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	response := &model.ResultAnnouncement{
		StatusCode: 200,
		Status:     true,
		Result:     announcements,
	}

	return response, nil
}

// ValidationPostAnnouncement validate input to create new announcement
func (a *AnnouncementUsecase) ValidationPostAnnouncement(ctx context.Context, ukmID int, input model.InputAnnouncement) (*model.SingleResultAnnouncement, *utils.Error) {
	if ukmID == 0 || &ukmID == nil {
		responseError := errors.New("ukm id can't be null or zero")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	if input.Title == "" || &input.Title == nil {
		responseError := errors.New("title can't be null or empty")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	if input.Description == "" || &input.Description == nil {
		responseError := errors.New("description can't be null or empty")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	if input.Time == "" || &input.Time == nil {
		responseError := errors.New("time can't be zero")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	announcement, err := a.AnnouncementRepo.Posting(ctx, ukmID, input)
	if err != nil {
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	response := &model.SingleResultAnnouncement{
		StatusCode: 200,
		Status:     true,
		Result:     announcement,
	}

	return response, nil
}

// UpdateOneAnnouncement update only one announcement
func (a *AnnouncementUsecase) UpdateOneAnnouncement(ctx context.Context, announcementID int, input model.InputAnnouncement) (*model.SingleResultAnnouncement, *utils.Error) {
	if announcementID == 0 || &announcementID == nil {
		responseError := errors.New("announcement id can't be null or zero")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	if input.Title == "" || &input.Title == nil {
		responseError := errors.New("title can't be null or empty")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	if input.Description == "" || &input.Description == nil {
		responseError := errors.New("description can't be null or empty")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	if input.Time == "" || &input.Time == nil {
		responseError := errors.New("time can't be null or empty")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	announcement, err := a.AnnouncementRepo.Update(ctx, announcementID, input)
	if err != nil {
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	response := &model.SingleResultAnnouncement{
		StatusCode: 200,
		Status:     true,
		Result:     announcement,
	}

	return response, nil
}

// DeleteOneAnnouncement delete only one announcement
func (a *AnnouncementUsecase) DeleteOneAnnouncement(ctx context.Context, announcementID int) (*model.SingleResultAnnouncement, *utils.Error) {
	if announcementID == 0 || &announcementID == nil {
		responseError := errors.New("announcement id can't be null or zero")
		return nil, utils.WrapErrorJson(responseError, fasthttp.StatusBadRequest)
	}

	err := a.AnnouncementRepo.Delete(ctx, announcementID)
	if err != nil {
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	response := &model.SingleResultAnnouncement{
		StatusCode: 200,
		Status:     true,
		Result:     nil,
	}

	return response, nil
}
