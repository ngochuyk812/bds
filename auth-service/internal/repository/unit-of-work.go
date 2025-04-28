package repository

import (
	repositorysite "auth_service/internal/repository/site"
	repositoryuser "auth_service/internal/repository/user"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type unitOfWork struct {
	db    *gorm.DB
	repos map[string]interface{}
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{
		db:    db,
		repos: make(map[string]interface{}),
	}
}

type UnitOfWork interface {
	GetSiteRepository() repositorysite.SiteRepository
	GetUserRepository() repositoryuser.UserRepository
	GetUserDetailRepository() repositoryuser.UserDetailRepository
	ExecTx(ctx context.Context, fn func(uow UnitOfWork) error) error
}

func (u *unitOfWork) ExecTx(ctx context.Context, fn func(uow UnitOfWork) error) error {
	tx := u.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	txUow := &unitOfWork{
		db:    tx,
		repos: make(map[string]interface{}),
	}

	err := fn(txUow)
	if err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit().Error
}

func (u *unitOfWork) GetSiteRepository() repositorysite.SiteRepository {
	const key = "SiteRepository"
	if repo, ok := u.repos[key]; ok {
		return repo.(repositorysite.SiteRepository)
	}
	repo := repositorysite.NewSiteRepository(u.db)
	u.repos[key] = repo
	return repo
}

func (u *unitOfWork) GetUserRepository() repositoryuser.UserRepository {
	const key = "UserRepository"
	if repo, ok := u.repos[key]; ok {
		return repo.(repositoryuser.UserRepository)
	}
	repo := repositoryuser.NewUserRepository(u.db)
	u.repos[key] = repo
	return repo
}

func (u *unitOfWork) GetUserDetailRepository() repositoryuser.UserDetailRepository {
	const key = "UserDetailRepository"
	if repo, ok := u.repos[key]; ok {
		return repo.(repositoryuser.UserDetailRepository)
	}
	repo := repositoryuser.NewUserDetailRepository(u.db)
	u.repos[key] = repo
	return repo
}
