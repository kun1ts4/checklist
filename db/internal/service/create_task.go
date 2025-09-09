package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kun1ts4/checklist/db/internal"
	"github.com/sirupsen/logrus"
)

func (s *TaskService) CreateTask(ctx context.Context, text string) (*internal.Task, error) {
	task := internal.Task{
		Id:   uuid.NewString(),
		Text: text,
	}

	err := s.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "create task error",
			"error":   err,
		}).Error()

		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"message": "create task success",
	})
	return &task, nil
}
