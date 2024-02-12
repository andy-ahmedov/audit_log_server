package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/andy-ahmedov/audit_log_server/internal/config"
	"github.com/andy-ahmedov/audit_log_server/internal/repository"
	"github.com/andy-ahmedov/audit_log_server/internal/server"
	"github.com/andy-ahmedov/audit_log_server/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.Db.Username,
		Password: cfg.Db.Password,
	})
	opts.ApplyURI(cfg.Db.Uri)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.Db.Database)

	auditRepo := repository.NewAudit(db)
	auditService := service.New(auditRepo)
	auditSrv := server.NewAuditService(auditService)
	srv := server.New(auditSrv)

	fmt.Println("SERVER STARTED")

	if err = srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
