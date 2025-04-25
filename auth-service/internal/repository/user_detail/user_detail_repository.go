package repositoryuserdetail

import (
	"auth_service/internal/db/user_detail"
	"auth_service/internal/entity"
	"context"
	"database/sql"
	"time"
)

type UserDetailRepository interface {
	CreateUserDetail(ctx context.Context, userDetail *entity.UserDetail) error
	GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*entity.UserDetail, error)
	UpdateUserDetail(ctx context.Context, userDetail *entity.UserDetail) error
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

func (u *userDetailRepository) CreateUserDetail(ctx context.Context, userDetail *entity.UserDetail) error {
	return u.writeQueries.CreateUserDetail(ctx, user_detail.CreateUserDetailParams{
		UserGuid:  userDetail.UserGuid,
		FirstName: sql.NullString{String: userDetail.FirstName, Valid: len(userDetail.FirstName) > 0},
		LastName:  sql.NullString{String: userDetail.LastName, Valid: len(userDetail.LastName) > 0},
		Phone:     sql.NullString{String: userDetail.Phone, Valid: len(userDetail.Phone) > 0},
		Avatar:    sql.NullString{String: userDetail.Avatar, Valid: len(userDetail.Avatar) > 0},
		Address:   sql.NullString{String: userDetail.Address, Valid: len(userDetail.Address) > 0},
		Createdat: userDetail.Createdat,
	})
}

func (u *userDetailRepository) GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*entity.UserDetail, error) {
	result, err := u.readQueries.GetUserDetailByUserGuid(ctx, userGuid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &entity.UserDetail{
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

func (u *userDetailRepository) UpdateUserDetail(ctx context.Context, userDetail *entity.UserDetail) error {
	return u.writeQueries.UpdateUserDetailByUserGuid(ctx, user_detail.UpdateUserDetailByUserGuidParams{
		FirstName: sql.NullString{String: userDetail.FirstName, Valid: len(userDetail.FirstName) > 0},
		LastName:  sql.NullString{String: userDetail.LastName, Valid: len(userDetail.LastName) > 0},
		Phone:     sql.NullString{String: userDetail.Phone, Valid: len(userDetail.Phone) > 0},
		Avatar:    sql.NullString{String: userDetail.Avatar, Valid: len(userDetail.Avatar) > 0},
		Address:   sql.NullString{String: userDetail.Address, Valid: len(userDetail.Address) > 0},
		UpdatedAt: userDetail.Updatedat,
		UserGuid:  userDetail.UserGuid,
	})
}

func (u *userDetailRepository) DeleteUserDetail(ctx context.Context, userGuid string) error {
	return u.writeQueries.DeleteUserDetailByUserGuid(ctx, user_detail.DeleteUserDetailByUserGuidParams{
		Deletedat: sql.NullInt64{Int64: time.Now().Unix()},
		UserGuid:  userGuid,
	})
}
