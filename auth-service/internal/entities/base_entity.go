package entities

type BaseEntity struct {
	ID        int32  `gorm:"primaryKey"`
	Guid      string `gorm:"unique"`
	SiteId    string `gorm:"column:site_id"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt *int64 `gorm:"column:deleted_at"`
}

func (a *BaseEntity) SetGuid(guid string) {
	a.Guid = guid
}

func (a *BaseEntity) SetCreatedAt(createdAt int64) {
	a.CreatedAt = createdAt
}

func (a *BaseEntity) SetUpdatedAt(updatedAt int64) {
	a.UpdatedAt = updatedAt
}

func (a *BaseEntity) SetDeletedAt(deletedAt *int64) {
	a.DeletedAt = deletedAt
}
