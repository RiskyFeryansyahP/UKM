package repository

import (
	"context"
	"github.com/confus1on/UKM/ent"
	announcementField "github.com/confus1on/UKM/ent/announcement"
	ukmField "github.com/confus1on/UKM/ent/ukm"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/announcement"
)

// AnnouncementRepository hold database client to connect into database
type AnnouncementRepository struct {
	DB *ent.Client
}

// NewAnnouncementRepository initiate announcement repository to connect into database client
// return RepositoryAnnouncement as interface
func NewAnnouncementRepository(db *ent.Client) announcement.RepositoryAnnouncement {
	return &AnnouncementRepository{DB: db}
}

// GetUkmAnnouncement get specification announcement based on ukm
// @param ukmID id of ukm you want take the announcement
func (a *AnnouncementRepository) GetUkmAnnouncement(ctx context.Context, ukmID int) ([]*ent.Announcement, error) {
	announcements, err := a.DB.Announcement.Query().
		Where(
			announcementField.HasOwnerAnnouncementWith(
				ukmField.IDEQ(ukmID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return announcements, nil
}

// Posting post a new announcement
func (a *AnnouncementRepository) Posting(ctx context.Context, ukmID int, input model.InputAnnouncement) (*ent.Announcement, error) {
	announcement, err := a.DB.Announcement.Create().
		SetTitle(input.Title).
		SetDescription(input.Description).
		SetTime(input.Time).
		SetOwnerAnnouncementID(ukmID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

// Update specific ukm
func (a *AnnouncementRepository) Update(ctx context.Context, announcementID int, input model.InputAnnouncement) (*ent.Announcement, error) {
	announcement, err := a.DB.Announcement.UpdateOneID(announcementID).
		SetTitle(input.Title).
		SetDescription(input.Description).
		SetTime(input.Time).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

// Delete one ukm and only return error
func (a *AnnouncementRepository) Delete(ctx context.Context, announcementID int) error {
	err := a.DB.Announcement.DeleteOneID(
		announcementID,
	).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
