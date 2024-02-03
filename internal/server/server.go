package server

import (
	"fmt"
	"net"

	audit "github.com/andy-ahmedov/audit_log_server/pkg/domain"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer  *grpc.Server
	auditServer audit.AuditServiceServer
}

func New(auditServer audit.AuditServiceServer) *Server {
	return &Server{
		grpcServer:  grpc.NewServer(),
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf("%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	audit.RegisterAuditServiceServer(s.grpcServer, s.auditServer)

	err = s.grpcServer.Serve(lis)

	return err
}
