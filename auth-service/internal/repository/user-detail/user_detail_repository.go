package repositoryuserdetail

import (
	"auth_service/internal/domain/user_detail"
	usercase "auth_service/internal/usecase/user"
	"context"
	"database/sql"
	"time"
)

type UserDetailRepository interface {
	CreateUserDetail(ctx context.Context, arg *usercase.CreateUserDetailUsecase) error
	GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*usercase.UserDetail, error)
	UpdateUserDetail(ctx context.Context, arg *usercase.UpdateUserDetailUsecase) error
	DeleteUserDetail(ctx context.Context, userGuid string) error
}

type userDetailRepository struct {
	readQueries  *user_detail.Queries
	writeQueries *user_detail.Queries
}

func NewUserDetailRepository(readDB, writeDB *sql.DB, tx *sql.Tx) UserDetailRepository {
	if tx != nil {
		txDB := user_detail.New(tx)
		return &userDetailRepository{
			readQueries:  txDB,
			writeQueries: txDB,
		}
	}
	return &userDetailRepository{
		readQueries:  user_detail.New(readDB),
		writeQueries: user_detail.New(writeDB),
	}
}

func (u *userDetailRepository) CreateUserDetail(ctx context.Context, arg *usercase.CreateUserDetailUsecase) error {
	return u.writeQueries.CreateUserDetail(ctx, user_detail.CreateUserDetailParams{
		UserGuid:  arg.UserGuid,
		FirstName: sql.NullString{String: arg.FirstName, Valid: len(arg.FirstName) > 0},
		LastName:  sql.NullString{String: arg.LastName, Valid: len(arg.LastName) > 0},
		Phone:     sql.NullString{String: arg.Phone, Valid: len(arg.Phone) > 0},
		Avatar:    sql.NullString{String: arg.Avatar, Valid: len(arg.Avatar) > 0},
		Address:   sql.NullString{String: arg.Address, Valid: len(arg.Address) > 0},
		Createdat: time.Now().Unix(),
	})
}

func (u *userDetailRepository) GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*usercase.UserDetail, error) {
	result, err := u.readQueries.GetUserDetailByUserGuid(ctx, userGuid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &usercase.UserDetail{
		ID:        result.ID,
		UserGuid:  result.UserGuid,
		FirstName: result.FirstName.String,
		LastName:  result.LastName.String,
		Phone:     result.Phone.String,
		Avatar:    result.Avatar.String,
		Address:   result.Address.String,
		Createdat: result.Createdat,
		Updatedat: result.Updatedat,
		Deletedat: result.Deletedat,
	}, err
}

func (u *userDetailRepository) UpdateUserDetail(ctx context.Context, arg *usercase.UpdateUserDetailUsecase) error {
	return u.writeQueries.UpdateUserDetailByUserGuid(ctx, user_detail.UpdateUserDetailByUserGuidParams{
		FirstName: sql.NullString{String: arg.FirstName, Valid: len(arg.FirstName) > 0},
		LastName:  sql.NullString{String: arg.LastName, Valid: len(arg.LastName) > 0},
		Phone:     sql.NullString{String: arg.Phone, Valid: len(arg.Phone) > 0},
		Avatar:    sql.NullString{String: arg.Avatar, Valid: len(arg.Avatar) > 0},
		Address:   sql.NullString{String: arg.Address, Valid: len(arg.Address) > 0},
		UpdatedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		UserGuid:  arg.UserGuid,
	})
}

func (u *userDetailRepository) DeleteUserDetail(ctx context.Context, userGuid string) error {
	return u.writeQueries.DeleteUserDetailByUserGuid(ctx, user_detail.DeleteUserDetailByUserGuidParams{
		Deletedat: sql.NullInt64{Int64: time.Now().Unix()},
		UserGuid:  userGuid,
	})
}
