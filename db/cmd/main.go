package main

import (
	"context"
	"github.com/kun1ts4/checklist/db/internal"
	"github.com/sirupsen/logrus"
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
}
