package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/kun1ts4/checklist/api/types"
	"github.com/labstack/echo/v4"
)

// CreateTask godoc
// @Summary      Создать задачу
// @Description  Создаёт новую задачу
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        request body types.CreateTaskRequest true "Данные задачи"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /create [post]
func CreateTask(c echo.Context) error {
	req := new(types.CreateTaskRequest)

	err := c.Bind(req)
	if err != nil {
		logrus.WithError(err).Warn("invalid request body")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//create task

	logrus.WithFields(logrus.Fields{
		"task_text": req.Text,
	}).Info("task created")
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "task created",
		"task":    req,
	})
}
