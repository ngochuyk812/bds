package entity

import (
	"database/sql"
	"time"
)

// UserDetail represents the user detail entity in the system
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

// NewUserDetail creates a new user detail entity
func NewUserDetail(userGuid, firstName, lastName, phone, avatar, address string) *UserDetail {
	return &UserDetail{
		UserGuid:  userGuid,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Avatar:    avatar,
		Address:   address,
		Createdat: time.Now().Unix(),
	}
}

// FullName returns the user's full name
func (ud *UserDetail) FullName() string {
	return ud.FirstName + " " + ud.LastName
}

// UpdateProfile updates the user's profile information
func (ud *UserDetail) UpdateProfile(firstName, lastName, phone, address string) {
	ud.FirstName = firstName
	ud.LastName = lastName
	ud.Phone = phone
	ud.Address = address
	ud.Updatedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}

// UpdateAvatar updates the user's avatar
func (ud *UserDetail) UpdateAvatar(avatar string) {
	ud.Avatar = avatar
	ud.Updatedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}

// MarkAsDeleted marks the user detail as deleted
func (ud *UserDetail) MarkAsDeleted() {
	ud.Deletedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}
