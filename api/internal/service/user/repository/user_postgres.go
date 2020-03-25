package repository

import (
	"time"

	"github.com/RiskyFeryansyahP/UKM/ent"
	userAgregate "github.com/RiskyFeryansyahP/UKM/ent/user"
	"github.com/RiskyFeryansyahP/UKM/internal/model"
	"github.com/RiskyFeryansyahP/UKM/internal/service/user"

	"github.com/valyala/fasthttp"
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
func (u *UserRepository) Register(ctx *fasthttp.RequestCtx, input model.InputCreateUser) (*ent.User, error) {
	now := time.Now()

	user, err := u.DB.User.Create().
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	_, err = u.DB.Profile.Create().
		SetFirstName(input.UserProfile.FirstName).
		SetLastName(input.UserProfile.LastName).
		SetYearGeneration(input.UserProfile.YearGeneration).
		SetPhone(input.UserProfile.Phone).
		SetStatus(true).
		SetImgProfile(input.UserProfile.ImgProfile).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		AddOwner(user).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login user signin with return profile user
func (u *UserRepository) Login(ctx *fasthttp.RequestCtx, input model.InputLoginUser) (*ent.Profile, error) {
	profile, err := u.DB.User.
		Query().
		Where(
			userAgregate.And(
				userAgregate.EmailEQ(input.Email),
				userAgregate.Password(input.Password),
			),
		).
		QueryProfile().
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return profile, nil
}
