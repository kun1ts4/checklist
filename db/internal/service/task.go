package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/kun1ts4/checklist/db/internal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

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

//func (s *TaskService) CompleteTask(ctx context.Context, task *internal.Task) (*internal.Task, error) {
//	logrus.WithFields(logrus.Fields{
//		"message": "complete task",
//		"task":    task,
//	}).Debug("complete task")
//	err := s.db.WithContext(ctx).Model(&internal.Task{}).Where("id = ?", task.Id).Updates(task).Error
//
//	logrus.WithFields(logrus.Fields{
//		"err": err,
//	}).Debug("complete task")
//
//	if err != nil {
//		logrus.WithFields(logrus.Fields{
//			"message": "complete task error",
//			"error":   err,
//		}).Error()
//	}
//
//	logrus.WithFields(logrus.Fields{
//		"message": "complete task success",
//	})
//	return task, nil
//}
