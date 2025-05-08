package repositoryuser

import (
	"auth_service/internal/entities"
	repositorybase "auth_service/internal/repository/base"
	"context"
	"errors"

	"github.com/ngochuyk812/building_block/infrastructure/helpers"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetBaseRepository() repositorybase.Repository[entities.User]
}

type userRepository struct {
	db   *gorm.DB
	base repositorybase.Repository[entities.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db:   db,
		base: repositorybase.NewRepository[entities.User](db),
	}
}

func (u *userRepository) GetBaseRepository() repositorybase.Repository[entities.User] {
	return u.base
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	authContext, ok := helpers.AuthContext(ctx)
	if !ok {
		return nil, errors.New("cannot get auth context")
	}

	var user entities.User
	err := u.db.WithContext(ctx).
		Where("email = ? AND site_id = ?", email, authContext.IdSite).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}
