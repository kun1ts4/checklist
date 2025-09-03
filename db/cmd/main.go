package main

import (
	"context"
	"github.com/kun1ts4/checklist/db/internal"
	server "github.com/kun1ts4/checklist/db/internal/grpc"
	"github.com/kun1ts4/checklist/db/internal/service"
	gen "github.com/kun1ts4/checklist/db/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	ctx := context.Background()
	config := internal.DBConfig{
		Host:     "postgres",
		Port:     "5432",
		User:     "user",
		Password: "password",
		Database: "postgres",
		SSLMode:  "disable",
	}

	db, err := internal.NewGorm(ctx, config)
	if err != nil {
		logrus.Fatal(err)
	}

	var one int
	if err := db.Raw("SELECT 1").Scan(&one).Error; err != nil {
		logrus.Fatalf("db check failed: %v", err)
	}
	logrus.Info("db connected")

	err = db.AutoMigrate(&internal.Task{})
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("db migrated")

	taskService := service.NewTaskService(db)
	taskServer := server.NewServer(taskService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	gen.RegisterTaskServiceServer(grpcServer, taskServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "grpc server failed",
			"error":   err,
		})
		return
	}
}
