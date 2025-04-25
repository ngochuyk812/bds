package repository

import (
	repositorysite "auth_service/internal/repository/site"
	repositoryuser "auth_service/internal/repository/user"
	repositoryuserdetail "auth_service/internal/repository/user_detail"
	"context"
	"database/sql"
	"fmt"
)

type unitOfWork struct {
	writeDB  *sql.DB
	readerDB *sql.DB
	tx       *sql.Tx
	repos    map[string]interface{}
}

func NewUnitOfWork(writeDB *sql.DB, readerDB *sql.DB) UnitOfWork {
	return newTxUnitOfWork(writeDB, readerDB, nil)
}
func newTxUnitOfWork(writeDB *sql.DB, readerDB *sql.DB, tx *sql.Tx) UnitOfWork {
	return &unitOfWork{
		writeDB:  writeDB,
		readerDB: readerDB,
		tx:       tx,
		repos:    make(map[string]interface{}),
	}
}

type UnitOfWork interface {
	GetSiteRepository() repositorysite.SiteRepository
	GetUserRepository() repositoryuser.UserRepository
	GetUserDetailRepository() repositoryuserdetail.UserDetailRepository

	ExecTx(ctx context.Context, fn func(uow UnitOfWork) error) error
}

func (u *unitOfWork) ExecTx(ctx context.Context, fn func(uow UnitOfWork) error) error {
	tx, err := u.writeDB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	txUow := newTxUnitOfWork(u.writeDB, u.writeDB, tx)
	err = fn(txUow)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (u *unitOfWork) GetSiteRepository() repositorysite.SiteRepository {
	const key = "SiteRepository"
	if repo, ok := u.repos[key]; ok {
		return repo.(repositorysite.SiteRepository)
	}
	siteRepo := repositorysite.NewSiteRepository(u.writeDB, u.readerDB, u.tx)
	u.repos[key] = siteRepo
	return siteRepo
}

func (u *unitOfWork) GetUserRepository() repositoryuser.UserRepository {
	const key = "UserRepository"
	if repo, ok := u.repos[key]; ok {
		return repo.(repositoryuser.UserRepository)
	}
	repo := repositoryuser.NewUserRepository(u.writeDB, u.readerDB, u.tx)
	u.repos[key] = repo
	return repo
}

func (u *unitOfWork) GetUserDetailRepository() repositoryuserdetail.UserDetailRepository {
	const key = "UserDetailRepository"
	if repo, ok := u.repos[key]; ok {
		return repo.(repositoryuserdetail.UserDetailRepository)
	}
	repo := repositoryuserdetail.NewUserDetailRepository(u.writeDB, u.readerDB, u.tx)
	u.repos[key] = repo
	return repo
}
