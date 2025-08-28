package handler

import (
	"net/http"

	"github.com/kun1ts4/checklist/api/types"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

// DeleteTask godoc
// @Summary      Удалить задачу
// @Description  Удаляет задачу по указанному ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        request body types.DeleteTaskRequest true "ID задачи для удаления"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /delete [delete]
func DeleteTask(c echo.Context) error {
	req := new(types.DeleteTaskRequest)
	if err := c.Bind(req); err != nil {
		logrus.WithError(err).Warn("invalid request body")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//delete task

	logrus.WithFields(logrus.Fields{
		"task_id": req.Id,
	}).Info("task deleted")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "task deleted",
		"task":    req,
	})
}
