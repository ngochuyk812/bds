package usercase

import "database/sql"

type GetUserByEmailUsecase struct {
	Email string
}

type CreateUserUsercase struct {
	Guid         string
	Email        string
	HashPassword string
	Salt         string
	LastName     string
	FirstName    string
}

type UserUsercase struct {
	ID           int32
	Guid         string
	Siteid       string
	Email        string
	HashPassword string
	Salt         string
	Createdat    int64
	Updatedat    sql.NullInt64
	Deletedat    sql.NullInt64
}
