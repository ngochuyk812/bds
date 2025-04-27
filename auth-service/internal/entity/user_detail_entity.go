package entity

import "database/sql"

type UserDetail struct {
	UserGuid  string
	FirstName sql.NullString
	LastName  sql.NullString
	Phone     sql.NullString
	Avatar    sql.NullString
	Address   sql.NullString
	BaseEntity
}

func (a *UserDetail) TableName() string {
	return "user_details"
}
