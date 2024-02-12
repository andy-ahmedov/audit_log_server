package server

import (
	"context"

	audit "github.com/andy-ahmedov/audit_log_server/pkg/domain"
)

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) error
}

type AuditServer struct {
	service AuditService
	audit.UnimplementedAuditServiceServer
}

func NewAuditService(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (as *AuditServer) Log(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	err := as.service.Insert(ctx, req)

	return &audit.Empty{}, err
}
