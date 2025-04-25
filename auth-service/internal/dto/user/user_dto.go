package userdto

import "database/sql"

type GetUserByEmailDTO struct {
	Email string
}

type UpdateUserDTO struct {
	Guid         string
	Email        string
	Active       bool
	HashPassword string
	Salt         string
}

type CreateUserDTO struct {
	Guid         string
	Email        string
	HashPassword string
	Salt         string
	LastName     string
	FirstName    string
}

type UserDTO struct {
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

type CreateUserDetailDTO struct {
	UserGuid  string
	FirstName string
	LastName  string
	Phone     string
	Avatar    string
	Address   string
}

type UpdateUserDetailDTO struct {
	UserGuid  string
	FirstName string
	LastName  string
	Phone     string
	Avatar    string
	Address   string
}

type UserDetailDTO struct {
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
