package location

import "github.com/labstack/echo/v4"

type Handlers interface {
	GetLocationBySatellites() echo.HandlerFunc
}
