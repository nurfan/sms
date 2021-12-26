package http

import (
	"context"
	"net/http"

	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	accidentsAct "github.com/nurfan/sms/action/accidents"
)

// FetchAccident for get list issue
func (adp *Adapter) FetchAccident(c echo.Context) error {
	ctx := context.Background()

	act := accidentsAct.NewFecthAccidents(adp.repoPsql)
	result, err := act.Handle(ctx)
	var resp Response

	if err != nil {
		resp.SetErrorResponse(err.Code, err.Message)
		return c.JSON(resp.Code, resp)
	}

	data := structs.Map(result)

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}

// FetchObjectCategory for get list issue
func (adp *Adapter) FetchObjectCategory(c echo.Context) error {
	ctx := context.Background()

	act := accidentsAct.NewFecthObject(adp.repoPsql)
	result, err := act.Handle(ctx)
	var resp Response

	if err != nil {
		resp.SetErrorResponse(err.Code, err.Message)
		return c.JSON(resp.Code, resp)
	}

	data := structs.Map(result)

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}
