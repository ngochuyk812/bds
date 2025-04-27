package entities

type UserDetail struct {
	UserGuid  string
	FirstName *string
	LastName  *string
	Phone     *string
	Avatar    *string
	Address   *string
	BaseEntity
}

func (a *UserDetail) TableName() string {
	return "user_details"
}
