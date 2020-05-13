package announcement

import (
	"context"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
)

// RepositoryAnnouncement hold interface value for repository
type RepositoryAnnouncement interface {
	GetUkmAnnouncement(ctx context.Context, ukmID int) ([]*ent.Announcement, error)
	Posting(ctx context.Context, ukmID int, input model.InputAnnouncement) (*ent.Announcement, error)
	Update(ctx context.Context, announcementID int, input model.InputAnnouncement) (*ent.Announcement, error)
	Delete(ctx context.Context, announcementID int) error
}
