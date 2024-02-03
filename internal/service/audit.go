package service

import (
	"context"

	audit "github.com/andy-ahmedov/audit_log_server/pkg/domain"
)

type Repository interface {
	Insert(ctx context.Context, req audit.LogItem) error
}

type Audit struct {
	repo Repository
}

func New(repo Repository) *Audit {
	return &Audit{
		repo: repo,
	}
}

func (a *Audit) Insert(ctx context.Context, req *audit.LogRequest) error {
	item := audit.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return a.repo.Insert(ctx, item)
}
