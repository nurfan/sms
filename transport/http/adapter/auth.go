package http

import (
	"context"
	"net/http"

	authAct "github.com/nurfan/sms/action/auth"

	"github.com/nurfan/sms/model"

	"github.com/labstack/echo/v4"
)

// GetToken get JWT token
func (adp *Adapter) GetToken(c echo.Context) error {
	ctx := context.Background()

	u := new(model.GetTokenRequest)
	if err := c.Bind(u); err != nil {
		return err
	}

	req := model.GetTokenRequest{
		Username: u.Username,
		Password: u.Password,
		AppKey:   c.Request().Header.Get("app-key"),
	}

	act := authAct.NewGetToken(adp.repoPsql)
	result, err := act.Handle(ctx, req)
	var resp Response

	if err != nil {
		resp.SetErrorResponse(int(err.Code), err.Message)
		return c.JSON(int(err.Code), resp)
	}

	data := map[string]interface{}{
		"Token": result.Data.Token,
	}

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}
