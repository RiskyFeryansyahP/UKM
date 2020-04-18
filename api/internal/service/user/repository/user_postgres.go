package repository

import (
	"context"
	"time"

	"github.com/confus1on/UKM/ent"
	userAgregate "github.com/confus1on/UKM/ent/user"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/user"
)

// UserRepository a struct to inject depedency from client connection
type UserRepository struct {
	DB *ent.Client
}

// NewUserRepository a construct initialize UserRepository
func NewUserRepository(DB *ent.Client) user.RepositoryUser {
	return &UserRepository{DB: DB}
}

// Register is a register user with return *ent.User and error
func (u *UserRepository) Register(ctx context.Context, input model.InputCreateUser) (*ent.User, error) {
	now := time.Now()

	tx, _ := u.DB.Tx(ctx)

	user, err := tx.User.Create().
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetRoleID(input.Role).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		return nil, rollback(tx, err)
	}

	_, err = tx.Profile.Create().
		SetFirstName(input.UserProfile.FirstName).
		SetLastName(input.UserProfile.LastName).
		SetJk(input.UserProfile.Jk).
		SetYearGeneration(input.UserProfile.YearGeneration).
		SetPhone(input.UserProfile.Phone).
		SetStatus(true).
		SetImgProfile(input.UserProfile.ImgProfile).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		AddOwner(user).
		Save(ctx)

	if err != nil {
		return nil, rollback(tx, err)
	}

	_ = tx.Commit()

	return user, nil
}

// Login user signin with return profile user
func (u *UserRepository) Login(ctx context.Context, input model.InputLoginUser) (*ent.User, error) {
	user, err := u.DB.User.
		Query().
		Where(
			userAgregate.And(
				userAgregate.EmailEQ(input.Email),
				userAgregate.Password(input.Password),
			),
		).
		WithProfile().
		WithRole().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func rollback(tx *ent.Tx, err error) error {
	_ = tx.Rollback() // doing rollback
	return err
}
