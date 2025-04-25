package entity

import (
	"database/sql"
	"time"
)

// Site represents the site entity in the system
type Site struct {
	ID        int32
	Guid      string
	Siteid    string
	Name      string
	Createdat int64
	Updatedat sql.NullInt64
	Deletedat sql.NullInt64
}

// NewSite creates a new site entity
func NewSite(siteid, name string) *Site {
	return &Site{
		Siteid:    siteid,
		Name:      name,
		Createdat: time.Now().Unix(),
	}
}

// IsDeleted checks if the site is deleted
func (s *Site) IsDeleted() bool {
	return s.Deletedat.Valid
}

// UpdateName updates the site's name
func (s *Site) UpdateName(name string) {
	s.Name = name
	s.Updatedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}

// UpdateSiteId updates the site's ID
func (s *Site) UpdateSiteId(siteid string) {
	s.Siteid = siteid
	s.Updatedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}

// MarkAsDeleted marks the site as deleted
func (s *Site) MarkAsDeleted() {
	s.Deletedat = sql.NullInt64{Int64: time.Now().Unix(), Valid: true}
}
