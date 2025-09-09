package server

import (
	"context"

	gen "github.com/kun1ts4/checklist/db/proto"
)

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
