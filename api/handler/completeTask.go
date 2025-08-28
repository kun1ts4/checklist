package handler

import (
	"net/http"

	"github.com/kun1ts4/checklist/api/types"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

// CompleteTask godoc
// @Summary      Завершить задачу
// @Description  Отмечает задачу как выполненную
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        request body types.CompleteTaskRequest true "ID задачи для завершения"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /done [put]
func CompleteTask(c echo.Context) error {
	req := new(types.CompleteTaskRequest)
	err := c.Bind(req)
	if err != nil {
		logrus.WithError(err).Warn("invalid request body")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//complete task

	logrus.WithFields(logrus.Fields{
		"task_id": req.Id,
	}).Info("task completed")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "task completed",
	})
}
