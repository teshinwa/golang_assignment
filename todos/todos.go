package todos

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/teshinwa/golang_assignment/logger"
)

func NewNewTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.AutoMigrate(Task{})

		var todo struct {
			Task string `json:"task"`
		}

		logger := logger.Extract(c)
		logger.Info("new task todo........")

		if err := c.Bind(&todo); err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": errors.Wrap(err, "new task").Error(),
			})
		}

		if err := db.Create(&Task{
			Task: todo.Task,
		}).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "create task").Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{})
	}
}

func GetTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//db.AutoMigrate(Task{})

		logger := logger.Extract(c)
		logger.Info("new task todo........")
		id := c.Param("id")
		var todo []Task
		if err := db.Find(&todo, id).Error; err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": errors.Wrap(err, "new task").Error(),
			})
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func UpdateTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//db.AutoMigrate(Task{})

		logger := logger.Extract(c)
		logger.Info("new task todo........ %")

		id := c.Param("id")

		var todo Task

		if err := db.Model(&todo).Where("id = ?", id).Update("processed", true).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "new task").Error(),
			})
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("%v has been updated", id))
	}
}

func DeleteTaskHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		//db.AutoMigrate(Task{})

		logger := logger.Extract(c)
		logger.Info("new task todo........ %")

		id := c.Param("id")

		var todo Task

		if err := db.Delete(&todo, id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"err": errors.Wrap(err, "new task").Error(),
			})
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("%v has been deleted", id))
	}
}

type Task struct {
	gorm.Model
	Task      string
	Processed bool
}

func (Task) TableName() string {
	return "todos"
}
