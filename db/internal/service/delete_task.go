package service

import (
	"context"

	"github.com/kun1ts4/checklist/db/internal"
	"github.com/sirupsen/logrus"
)

func (s *TaskService) DeleteTask(ctx context.Context, task *internal.Task) (*internal.Task, error) {
	err := s.db.WithContext(ctx).Model(&internal.Task{}).Where("id = ?", task.Id).Delete(&internal.Task{}).Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": "delete task error",
			"error":   err,
		}).Error()
	}

	logrus.WithFields(logrus.Fields{
		"message": "delete task success",
	})
	return task, nil
}
