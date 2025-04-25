package repositoryuser

import (
	"auth_service/internal/db/user"
	"auth_service/internal/entity"
	"context"
	"database/sql"
	"errors"

	"github.com/ngochuyk812/building_block/infrastructure/helpers"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByGuid(ctx context.Context, guid string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
}

type userRepository struct {
	readQueries  *user.Queries
	writeQueries *user.Queries
}

func (u *userRepository) CreateUser(ctx context.Context, userEntity *entity.User) error {
	authContext, oke := helpers.AuthContext(ctx)
	if !oke {
		return errors.New("cannot get auth context")
	}

	err := u.readQueries.CreateUser(ctx, user.CreateUserParams{
		Email:        userEntity.Email,
		Siteid:       authContext.IdSite,
		Guid:         userEntity.Guid,
		HashPassword: userEntity.HashPassword,
		Salt:         userEntity.Salt,
		Createdat:    userEntity.Createdat,
	})

	return err
}

func (u *userRepository) UpdateUser(ctx context.Context, userEntity *entity.User) error {
	err := u.readQueries.UpdateUserByGuid(ctx, user.UpdateUserByGuidParams{
		Email: sql.NullString{
			String: userEntity.Email,
			Valid:  len(userEntity.Email) > 0,
		},
		Guid:         userEntity.Guid,
		UpdatedAt:    userEntity.Updatedat,
		Active:       userEntity.Active,
		HashPassword: sql.NullString{String: userEntity.HashPassword, Valid: len(userEntity.HashPassword) > 0},
		Salt:         sql.NullString{String: userEntity.Salt, Valid: len(userEntity.Salt) > 0},
	})

	return err
}
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	authContext, oke := helpers.AuthContext(ctx)
	if !oke {
		return nil, errors.New("cannot get auth context")
	}
	rs, err := u.readQueries.GetUserByEmail(ctx, user.GetUserByEmailParams{
		Email:  email,
		Siteid: authContext.IdSite,
	})
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &entity.User{
		ID:           rs.ID,
		Guid:         rs.Guid,
		Siteid:       rs.Siteid,
		Email:        rs.Email,
		HashPassword: rs.HashPassword,
		Salt:         rs.Salt,
		Active:       rs.Active,
		Createdat:    rs.Createdat,
		Updatedat:    rs.Updatedat,
		Deletedat:    rs.Deletedat,
	}, err
}

func (u *userRepository) GetUserByGuid(ctx context.Context, guid string) (*entity.User, error) {
	rs, err := u.readQueries.GetUserByGuid(ctx, guid)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &entity.User{
		ID:           rs.ID,
		Guid:         rs.Guid,
		Siteid:       rs.Siteid,
		Email:        rs.Email,
		HashPassword: rs.HashPassword,
		Salt:         rs.Salt,
		Active:       rs.Active,
		Createdat:    rs.Createdat,
		Updatedat:    rs.Updatedat,
		Deletedat:    rs.Deletedat,
	}, err
}
func NewUserRepository(readDB, writeDB *sql.DB, tx *sql.Tx) UserRepository {

	if tx != nil {
		txDB := user.New(tx)
		return &userRepository{
			readQueries:  txDB,
			writeQueries: txDB,
		}
	}
	return &userRepository{
		readQueries:  user.New(readDB),
		writeQueries: user.New(writeDB),
	}

}
