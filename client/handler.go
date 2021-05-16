package client

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rkiminius/carbon-based-life-forms/manager"
	"net/http"
)

func InitRouter(group *echo.Group) {
	group.GET("/minerals", askMineralsHandler)
	group.POST("/order", orderActionOnMineralHandler)
}

func askMineralsHandler(c echo.Context) error {
	minerals, _ := manager.GetAvailableMinerals()
	return c.JSON(http.StatusOK, minerals)
}

func orderActionOnMineralHandler(c echo.Context) error {
	var clientRequest ClientRequest
	err := json.NewDecoder(c.Request().Body).Decode(&clientRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request!")
	}

	if ok := clientRequest.Action.IsValid(); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid action type!")
	}

	err = PerformActionsOnMinerals(clientRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.String(http.StatusOK, "order placed successfully")
}
