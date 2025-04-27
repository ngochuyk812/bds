package repositoryuserdetail

import (
	"auth_service/internal/entities"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserDetailRepository interface {
	CreateUserDetail(ctx context.Context, userDetail *entities.UserDetail) error
	GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*entities.UserDetail, error)
	UpdateUserDetail(ctx context.Context, userDetail *entities.UserDetail) error
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

func (u *userDetailRepository) CreateUserDetail(ctx context.Context, userDetail *entities.UserDetail) error {
	now := time.Now().Unix()
	userDetail.CreatedAt = now
	return u.db.WithContext(ctx).Create(userDetail).Error
}

func (u *userDetailRepository) GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*entities.UserDetail, error) {
	var detail entities.UserDetail
	err := u.db.WithContext(ctx).Where("user_guid = ?", userGuid).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &detail, err
}

func (u *userDetailRepository) UpdateUserDetail(ctx context.Context, userDetail *entities.UserDetail) error {
	now := time.Now().Unix()
	userDetail.UpdatedAt = now

	return u.db.WithContext(ctx).
		Model(&entities.UserDetail{}).
		Where("user_guid = ?", userDetail.UserGuid).
		Updates(userDetail).Error
}

func (u *userDetailRepository) DeleteUserDetail(ctx context.Context, userGuid string) error {
	now := time.Now().Unix()
	return u.db.WithContext(ctx).
		Model(&entities.UserDetail{}).
		Where("user_guid = ?", userGuid).
		Update("deletedat", &now).Error
}
