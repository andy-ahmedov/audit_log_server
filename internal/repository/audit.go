package repository

import (
	"context"

	audit "github.com/andy-ahmedov/audit_log_server/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// создаем структуру аудит с единственным полем бд типа *mongo.Database
// создаем функцию NewAudit, которая возвращает эту структуру
// создаем метод Insert, принимающий контекст и audit.LogItem
// который вставляет второй параметр в Коллекцию логов с помощью дополнительного метода InsertOne, в который и передается контекст с самим логоm
// возвращаем ошибку

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
