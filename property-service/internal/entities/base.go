package entities

import (
	"context"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/ngochuyk812/building_block/infrastructure/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseEntity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Guid      string             `json:"guid" bson:"guid"`
	SiteID    string             `json:"site_id" bson:"site_id"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
	DeletedAt *int64             `json:"deleted_at" bson:"deleted_at"`
}

func NewBaseEntity(ctx context.Context) BaseEntity {
	auth, ok := helpers.AuthContext(ctx)
	siteId := ""
	if ok {
		siteId = auth.IdSite
	} else {
		siteId = "-1"
	}
	guid, _ := uuid.NewV4()
	now := time.Now().Unix()
	return BaseEntity{
		CreatedAt: now,
		UpdatedAt: now,
		Guid:      guid.String(),
		SiteID:    siteId,
	}
}
