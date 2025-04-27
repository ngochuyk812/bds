package entity

type Site struct {
	Name string
	BaseEntity
}

func (a *Site) TableName() string {
	return "sites"
}
