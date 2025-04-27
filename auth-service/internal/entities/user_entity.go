package entities

type User struct {
	Email        string
	HashPassword string
	Salt         string
	Active       bool
	BaseEntity
}

func (a *User) TableName() string {
	return "users"
}
