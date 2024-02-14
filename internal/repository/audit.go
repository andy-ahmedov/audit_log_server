package repository

import (
	"context"

	audit "github.com/andy-ahmedov/audit_log_server/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Audit struct {
	db *mongo.Database
}

func NewAudit(db *mongo.Database) *Audit {
	return &Audit{
		db: db,
	}
}

func (a *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	_, err := a.db.Collection("log").InsertOne(ctx, item)

	return err
}
