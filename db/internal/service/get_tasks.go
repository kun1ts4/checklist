package service

import (
	"context"

	"github.com/kun1ts4/checklist/db/internal"
	"github.com/sirupsen/logrus"
)

func (s *TaskService) GetTasks(ctx context.Context) ([]*internal.Task, error) {
	var tasks []*internal.Task

	err := s.db.WithContext(ctx).Find(&tasks).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "get tasks error",
			"error":   err,
		}).Error()
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"message": "get tasks success",
	})

	return tasks, nil
}
