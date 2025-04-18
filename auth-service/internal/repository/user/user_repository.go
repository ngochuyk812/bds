package repositoryuser

import (
	"auth_service/internal/domain/user"
	usercase "auth_service/internal/usecase/user"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ngochuyk812/building_block/infrastructure/helpers"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, arg usercase.GetUserByEmailUsecase) (*user.User, error)
	CreateUser(ctx context.Context, arg *usercase.CreateUserUsercase) error
}

type userRepository struct {
	readQueries  *user.Queries
	writeQueries *user.Queries
}

func (u *userRepository) CreateUser(ctx context.Context, arg *usercase.CreateUserUsercase) error {
	authContext, oke := helpers.AuthContext(ctx)
	if !oke {
		return errors.New("cannot get auth context")
	}

	err := u.readQueries.CreateUser(ctx, user.CreateUserParams{
		Email:        arg.Email,
		Siteid:       authContext.IdSite,
		Guid:         arg.Guid,
		HashPassword: arg.HashPassword,
		Salt:         arg.Salt,
		Createdat:    time.Now().Unix(),
	})

	return err
}
func (u *userRepository) GetUserByEmail(ctx context.Context, arg usercase.GetUserByEmailUsecase) (*user.User, error) {
	authContext, oke := helpers.AuthContext(ctx)
	if !oke {
		return nil, errors.New("cannot get auth context")
	}
	rs, err := u.readQueries.GetUserByEmail(ctx, user.GetUserByEmailParams{
		Email:  arg.Email,
		Siteid: authContext.IdSite,
	})

	return &rs, err
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
