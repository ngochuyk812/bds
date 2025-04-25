package entity

import (
	"database/sql"
	"time"
)

// User represents the user entity in the system
type User struct {
	ID           int32
	Guid         string
	Siteid       string
	Email        string
	HashPassword string
	Salt         string
	Active       sql.NullBool
	Createdat    int64
	Updatedat    sql.NullInt64
	Deletedat    sql.NullInt64
}

// NewUser creates a new user entity
func NewUser(guid, siteid, email, hashPassword, salt string) *User {
	return &User{
		Guid:         guid,
		Siteid:       siteid,
		Email:        email,
		HashPassword: hashPassword,
		Salt:         salt,
		Active:       sql.NullBool{Bool: false, Valid: true},
		Createdat:    time.Now().Unix(),
	}
}

// IsActive checks if the user is active
func (u *User) IsActive() bool {
	return u.Active.Valid && u.Active.Bool
}

// SetActive sets the user's active status
func (u *User) SetActive(active bool) {
	u.Active = sql.NullBool{Bool: active, Valid: true}
	u.Updatedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}

// UpdatePassword updates the user's password
func (u *User) UpdatePassword(hashPassword, salt string) {
	u.HashPassword = hashPassword
	u.Salt = salt
	u.Updatedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}

// UpdateEmail updates the user's email
func (u *User) UpdateEmail(email string) {
	u.Email = email
	u.Updatedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}

// MarkAsDeleted marks the user as deleted
func (u *User) MarkAsDeleted() {
	u.Deletedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}
