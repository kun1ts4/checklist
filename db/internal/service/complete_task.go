package service

//
//import (
//	"context"
//	"github.com/kun1ts4/checklist/db/internal"
//)
//
//func (s *TaskService) CompleteTask(ctx context.Context, id string) (*internal.Task, error) {
//	s.db.WithContext(ctx).Model(&internal.Task{}).Where("id = ?", id).Update("is_done", true)
//}
