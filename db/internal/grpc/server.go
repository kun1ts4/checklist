package server

import (
	"context"
	"github.com/kun1ts4/checklist/db/internal"
	"github.com/kun1ts4/checklist/db/internal/service"
	gen "github.com/kun1ts4/checklist/db/proto"
)

type Server struct {
	gen.UnimplementedTaskServiceServer
	taskService *service.TaskService
}

func NewServer(taskService *service.TaskService) *Server {
	return &Server{taskService: taskService}
}

func (s *Server) CreateTask(ctx context.Context, request *gen.CreateTaskRequest) (*gen.TaskResponse, error) {
	task, err := s.taskService.CreateTask(ctx, request.Text)
	if err != nil {
		return nil, err
	}

	return &gen.TaskResponse{
		Id:     task.Id,
		Text:   task.Text,
		IsDone: task.IsDone,
	}, nil
}

func (s *Server) GetTasks(ctx context.Context, request *gen.GetTaskRequest) (*gen.GetTasksResponse, error) {
	tasks, err := s.taskService.GetTasks(ctx)
	if err != nil {
		return nil, err
	}

	tasksResponse := &gen.GetTasksResponse{}
	for _, task := range tasks {
		tasksResponse.Tasks = append(tasksResponse.Tasks, &gen.TaskResponse{
			Id:     task.Id,
			Text:   task.Text,
			IsDone: task.IsDone,
		})
	}

	return tasksResponse, nil
}

//func (s *Server) CompleteTask(ctx context.Context, request *gen.CompleteTaskRequest) (*gen.TaskResponse, error) {
//	taskId := request.Id
//
//	task := internal.Task{Id: taskId}
//
//	logrus.WithFields(logrus.Fields{
//		"taskId": taskId,
//	}).Debug("complete task")
//
//	res, err := s.taskService.CompleteTask(ctx, &task)
//	if err != nil {
//		return nil, err
//	}
//
//	logrus.WithFields(logrus.Fields{
//		"res": res,
//	}).Debug("complete task")
//
//	return &gen.TaskResponse{Id: taskId, IsDone: true}, nil
//}

func (s *Server) DeleteTask(ctx context.Context, request *gen.DeleteTaskRequest) (*gen.DeleteTaskResponse, error) {
	taskId := request.Id
	task := &internal.Task{Id: taskId}

	_, err := s.taskService.DeleteTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return &gen.DeleteTaskResponse{Success: true}, nil
}
