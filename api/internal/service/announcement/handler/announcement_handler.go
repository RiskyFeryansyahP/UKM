package handler

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/announcement"

	"github.com/valyala/fasthttp"
)

// AnnouncementHandler hold announcement usecase value
type AnnouncementHandler struct {
	AnnouncementUC announcement.UsecaseAnnouncement
}

// NewAnnouncementHandler initiate announcement handler to inject announcement usecase
func NewAnnouncementHandler(announcementUC announcement.UsecaseAnnouncement) *AnnouncementHandler {
	return &AnnouncementHandler{AnnouncementUC: announcementUC}
}

// GetAnnouncement get announcement and return it as json
func (a *AnnouncementHandler) GetAnnouncement(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("ukmID").(string)

	ukmID, _ := strconv.Atoi(param)

	ctx.Response.Header.SetContentType("application/json")

	result, err := a.AnnouncementUC.GetSpecificAnnouncement(context.Background(), ukmID)
	if err != nil {
		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}

// PostAnnouncement post a new announcement
func (a *AnnouncementHandler) PostAnnouncement(ctx *fasthttp.RequestCtx) {
	var input model.InputAnnouncement

	param := ctx.UserValue("ukmID").(string)

	ukmID, _ := strconv.Atoi(param)

	body := ctx.Request.Body()
	_ = json.Unmarshal(body, &input)

	ctx.Response.Header.SetContentType("application/json")

	result, err := a.AnnouncementUC.ValidationPostAnnouncement(context.Background(), ukmID, input)
	if err != nil {
		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}

// UpdateAnnouncement update announcement and return it as json
func (a *AnnouncementHandler) UpdateAnnouncement(ctx *fasthttp.RequestCtx) {
	var input model.InputAnnouncement

	param := ctx.UserValue("announcementID").(string)

	annoncementID, _ := strconv.Atoi(param)

	body := ctx.Request.Body()
	_ = json.Unmarshal(body, &input)

	ctx.Response.Header.SetContentType("application/json")

	result, err := a.AnnouncementUC.UpdateOneAnnouncement(context.Background(), annoncementID, input)
	if err != nil {
		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}

// DeleteAnnouncement delete announcement
func (a *AnnouncementHandler) DeleteAnnouncement(ctx *fasthttp.RequestCtx) {
	param := ctx.UserValue("announcementID").(string)

	annoncementID, _ := strconv.Atoi(param)

	ctx.Response.Header.SetContentType("application/json")

	result, err := a.AnnouncementUC.DeleteOneAnnouncement(context.Background(), annoncementID)
	if err != nil {
		_ = json.NewEncoder(ctx).Encode(err)
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}
