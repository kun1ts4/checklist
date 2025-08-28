package handler

import (
	"net/http"

	"github.com/kun1ts4/checklist/api/types"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// GetTasks godoc
// @Summary      Получить список задач
// @Description  Возвращает список всех задач
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /list [get]
func GetTasks(c echo.Context) error {

	// get tasks

	logrus.Info("tasks requested")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": types.TasksList{Tasks: []types.Task{{}}},
	})
}
