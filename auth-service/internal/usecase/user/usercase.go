package usercase

import "database/sql"

type GetUserByEmailUsecase struct {
	Email string
}

type UpdateUserUsercase struct {
	Guid         string
	Email        string
	Active       bool
	HashPassword string
	Salt         string
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

type CreateUserDetailUsecase struct {
	UserGuid  string
	FirstName string
	LastName  string
	Phone     string
	Avatar    string
	Address   string
}

type UpdateUserDetailUsecase struct {
	UserGuid  string
	FirstName string
	LastName  string
	Phone     string
	Avatar    string
	Address   string
}

type UserDetail struct {
	ID        int32
	UserGuid  string
	FirstName string
	LastName  string
	Phone     string
	Avatar    string
	Address   string
	Createdat int64
	Updatedat sql.NullInt64
	Deletedat sql.NullInt64
}
