package api // 1

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(fasthttp.StatusOK, "Most Popular") // 2
	}
}