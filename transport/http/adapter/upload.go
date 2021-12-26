package http

import (
	"context"
	"net/http"

	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"

	uploadAct "github.com/nurfan/sms/action/upload"
)

// Upload for get list issue
func (adp *Adapter) Upload(c echo.Context) error {
	ctx := context.Background()

	file, err := c.FormFile("file")

	if err != nil {
		var resp Response
		errRes := adp.e.BadRequest(err.Error())
		resp.SetErrorResponse(errRes.Code, errRes.Message)
		return c.JSON(resp.Code, resp)
	}

	act := uploadAct.NewUpload()
	result, notok := act.Handle(ctx, file)
	var resp Response

	if notok != nil {
		resp.SetErrorResponse(notok.Code, notok.Message)
		return c.JSON(resp.Code, resp)
	}

	data := structs.Map(result)

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}
