package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/rkiminius/carbon-based-life-forms/client"
	"github.com/rkiminius/carbon-based-life-forms/config"
	"github.com/rkiminius/carbon-based-life-forms/manager"
	"net/http"
)

func init() {
	config.GetConfig("conf.yaml")
}

func main() {

	client.InitAmqp()

	e := echo.New()

	e.GET("/minerals", func(c echo.Context) error {
		minerals, _ := manager.GetAvailableMinerals()
		return c.JSON(http.StatusOK, minerals)
	})

	e.POST("/order", func(c echo.Context) error {
		var clientRequest client.ClientRequest
		err := json.NewDecoder(c.Request().Body).Decode(&clientRequest)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid request!")
		}

		if ok := clientRequest.Action.IsValid(); !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid action type!")
		}

		err = client.PerformActionsOnMinerals(clientRequest)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return c.String(http.StatusOK, "order placed successfully")
	})

	e.Logger.Fatal(e.Start(config.Conf.ClientPort))
}
