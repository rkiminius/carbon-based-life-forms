package manager

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	"github.com/rkiminius/carbon-based-life-forms/task"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func InitRouter(group *echo.Group) {
	group.GET("/mineralType/all", getAllMineralTypes)
	group.POST("/mineralType/new", postMineralType)
	group.DELETE("/mineralType/:id", deleteMineralType)
	group.GET("/task", getAllTasks)
	group.GET("/task/:taskId", getTask)
}

func getAllMineralTypes(c echo.Context) error {

	mineralTypeList, err := mineral.GetMineralTypeList()
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error on data retrieval.")
	}

	return c.JSON(http.StatusOK, mineralTypeList)
}

func postMineralType(c echo.Context) error {
	var mineralT mineral.MineralType

	err := json.NewDecoder(c.Request().Body).Decode(&mineralT)
	if err != nil {
		if !strings.Contains(err.Error(), "json: invalid use of ,string struct tag, trying to unmarshal") {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	mineralTypeFromDb, err := mineral.InsertMineralType(&mineralT)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, mineralTypeFromDb)
}

func deleteMineralType(c echo.Context) error {

	mineralTypeId := c.Param("id")

	deleted, err := mineral.DeleteItem(mineralTypeId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, deleted > 0)
}

func getAllTasks(c echo.Context) error {

	tasks, err := task.GetList()
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error on data retrieval.")
	}

	return c.JSON(http.StatusOK, tasks)
}

func getTask(c echo.Context) error {

	taskID := c.Param("taskId")

	objId, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error on data retrieval.")
	}

	taskFromDb, err := task.GetById(objId)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, taskFromDb)
}
