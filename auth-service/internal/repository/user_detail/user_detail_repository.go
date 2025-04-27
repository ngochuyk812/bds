package repositoryuserdetail

import (
	"auth_service/internal/entity"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserDetailRepository interface {
	CreateUserDetail(ctx context.Context, userDetail *entity.UserDetail) error
	GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*entity.UserDetail, error)
	UpdateUserDetail(ctx context.Context, userDetail *entity.UserDetail) error
	DeleteUserDetail(ctx context.Context, userGuid string) error
}

type userDetailRepository struct {
	db *gorm.DB
}

func NewUserDetailRepository(db *gorm.DB) UserDetailRepository {
	return &userDetailRepository{
		db: db,
	}
}

func (u *userDetailRepository) CreateUserDetail(ctx context.Context, userDetail *entity.UserDetail) error {
	now := time.Now().Unix()
	userDetail.Createdat = now
	return u.db.WithContext(ctx).Create(userDetail).Error
}

func (u *userDetailRepository) GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*entity.UserDetail, error) {
	var detail entity.UserDetail
	err := u.db.WithContext(ctx).Where("user_guid = ?", userGuid).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &detail, err
}

func (u *userDetailRepository) UpdateUserDetail(ctx context.Context, userDetail *entity.UserDetail) error {
	now := time.Now().Unix()
	userDetail.UpdatedAt = now

	return u.db.WithContext(ctx).
		Model(&entity.UserDetail{}).
		Where("user_guid = ?", userDetail.UserGuid).
		Updates(userDetail).Error
}

func (u *userDetailRepository) DeleteUserDetail(ctx context.Context, userGuid string) error {
	now := time.Now().Unix()
	return u.db.WithContext(ctx).
		Model(&entity.UserDetail{}).
		Where("user_guid = ?", userGuid).
		Update("deletedat", &now).Error
}
