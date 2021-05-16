package manager

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rkiminius/carbon-based-life-forms/mineral"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func InitRouter(group *echo.Group) {
	group.GET("/mineralType/all", getAllMineralTypes)
	group.POST("/mineralType/new", postMineralType)
	group.DELETE("/mineralType/:id", deleteMineralType)
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
