package siteusecase

import (
	"errors"
)

var (
	ErrSiteNotFound = errors.New("site_not_found")
	ErrSiteExist    = errors.New("site_exist")
)
