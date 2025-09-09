package server

import (
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
