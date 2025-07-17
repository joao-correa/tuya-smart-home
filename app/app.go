package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func BindApiRoutes(e *echo.Echo) {
	api, _, _ := LoadApis()

	e.POST("/api/apply-scene", func(c echo.Context) error {
		err := api.tuyaDeviceApi.ApplyScene(c.QueryParam("scene"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, "ok")
	})

	e.POST("/api/test", func(c echo.Context) error {
		test(api)
		return c.JSON(200, "ok")
	})
}

func test(api *Api) {
	err := api.tuyaDeviceApi.ApplyScene("low-warm")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(time.Duration(time.Second * 5))

	err = api.tuyaDeviceApi.ApplyScene("turn-off")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(time.Duration(time.Second * 5))
	err = api.tuyaDeviceApi.ApplyScene("turn-on")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(time.Duration(time.Second * 5))
	err = api.tuyaDeviceApi.ApplyScene("full-warm")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(time.Duration(time.Second * 5))
	err = api.tuyaDeviceApi.ApplyScene("mid-warm")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(time.Duration(time.Second * 5))
	err = api.tuyaDeviceApi.ApplyScene("low-warm")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(time.Duration(time.Second * 5))
	err = api.tuyaDeviceApi.ApplyScene("low-cold")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(time.Duration(time.Second * 5))
	err = api.tuyaDeviceApi.ApplyScene("mid-cold")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	time.Sleep(time.Duration(time.Second * 5))
	err = api.tuyaDeviceApi.ApplyScene("full-cold")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
