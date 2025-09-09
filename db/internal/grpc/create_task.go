package server

import (
	"context"

	gen "github.com/kun1ts4/checklist/db/proto"
)

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
