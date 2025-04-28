package repositoryuser

import (
	"auth_service/internal/entities"
	repositorybase "auth_service/internal/repository/base"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserDetailRepository interface {
	GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*entities.UserDetail, error)
	GetBaseRepository() repositorybase.Repository[entities.UserDetail]
}

type userDetailRepository struct {
	db   *gorm.DB
	base repositorybase.Repository[entities.UserDetail]
}

func NewUserDetailRepository(db *gorm.DB) UserDetailRepository {
	return &userDetailRepository{
		db:   db,
		base: repositorybase.NewRepository[entities.UserDetail](db),
	}
}

func (u *userDetailRepository) GetBaseRepository() repositorybase.Repository[entities.UserDetail] {
	return u.base
}

func (u *userDetailRepository) GetUserDetailByUserGuid(ctx context.Context, userGuid string) (*entities.UserDetail, error) {
	var detail entities.UserDetail
	err := u.db.WithContext(ctx).Where("user_guid = ?", userGuid).First(&detail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &detail, err
}

func (u *userDetailRepository) DeleteUserDetail(ctx context.Context, userGuid string) error {
	now := time.Now().Unix()
	return u.db.WithContext(ctx).
		Model(&entities.UserDetail{}).
		Where("user_guid = ?", userGuid).
		Update("deletedat", &now).Error
}
