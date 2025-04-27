package entities

type BaseEntity struct {
	ID        int32  `gorm:"primaryKey"`
	Guid      string `gorm:"unique"`
	SiteId    string
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt *int64 `gorm:"column:deleted_at"`
}
