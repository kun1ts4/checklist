package server

import (
	"context"

	"github.com/kun1ts4/checklist/db/internal"
	gen "github.com/kun1ts4/checklist/db/proto"
)

func (s *Server) DeleteTask(ctx context.Context, request *gen.DeleteTaskRequest) (*gen.DeleteTaskResponse, error) {
	taskId := request.Id
	task := &internal.Task{Id: taskId}

	_, err := s.taskService.DeleteTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return &gen.DeleteTaskResponse{Success: true}, nil
}
