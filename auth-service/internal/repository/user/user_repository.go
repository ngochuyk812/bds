package repositoryuser

import (
	"auth_service/internal/entities"
	"context"
	"errors"
	"time"

	"github.com/ngochuyk812/building_block/infrastructure/helpers"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByGuid(ctx context.Context, guid string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
	UpdateUser(ctx context.Context, user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, userEntity *entities.User) error {
	authContext, ok := helpers.AuthContext(ctx)
	if !ok {
		return errors.New("cannot get auth context")
	}

	userEntity.SiteId = authContext.IdSite
	userEntity.CreatedAt = time.Now().Unix()

	return u.db.WithContext(ctx).Create(userEntity).Error
}

func (u *userRepository) UpdateUser(ctx context.Context, userEntity *entities.User) error {
	userEntity.UpdatedAt = time.Now().Unix()

	return u.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("guid = ?", userEntity.Guid).
		Updates(userEntity).Error
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	authContext, ok := helpers.AuthContext(ctx)
	if !ok {
		return nil, errors.New("cannot get auth context")
	}

	var user entities.User
	err := u.db.WithContext(ctx).
		Where("email = ? AND siteid = ?", email, authContext.IdSite).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (u *userRepository) GetUserByGuid(ctx context.Context, guid string) (*entities.User, error) {
	var user entities.User
	err := u.db.WithContext(ctx).
		Where("guid = ?", guid).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}
