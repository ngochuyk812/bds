package db_helper

import (
	"context"

	"github.com/ngochuyk812/building_block/infrastructure/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func BuildFilter(ctx context.Context, customFilter bson.M) bson.M {
	merged := bson.M{"deleted_at": nil}
	authContext, ok := helpers.AuthContext(ctx)
	if ok {
		merged["site_id"] = authContext.IdSite
	}
	for k, v := range customFilter {
		merged[k] = v
	}
	return merged
}
