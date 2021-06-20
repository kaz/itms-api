package server

import (
	"fmt"
	"net/http"

	"github.com/kaz/itms-api/internal/itms"
	"github.com/labstack/echo/v4"
)

type (
	PatchSwitchRequest struct {
		Instrument itms.Instrument `json:"inst"`
		Mode       itms.Mode       `json:"mode"`
	}
)

func patchSwitch(c echo.Context) error {
	params := &PatchSwitchRequest{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("failed to parse request: %v", err))
	}
	if params.Instrument == "" {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("`inst` is required"))
	}
	if params.Mode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("`mode` is required"))
	}

	if err := itms.SwitchInstrumentMode(params.Instrument, params.Mode); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to switch instrument mode: %v", err))
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}
