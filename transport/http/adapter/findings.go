package http

import (
	"context"
	"net/http"

	"github.com/fatih/structs"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nurfan/sms/model"

	findingsAct "github.com/nurfan/sms/action/findings"
)

// GetFindings for get list issue
func (adp *Adapter) GetFindings(c echo.Context) error {
	ctx := context.Background()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	req := model.GetFindingsRequest{
		UserID:    claims["username"].(int32),
		RoleID:    claims["role_id"].(int32),
		SectionID: claims["section_id"].(int32),
		Page:      1,
		Limit:     15,
	}

	act := findingsAct.NewGetFindings(adp.repoPsql)
	result, err := act.Handle(ctx, req)
	var resp Response

	if err != nil {
		resp.SetErrorResponse(err.Code, err.Message)
		return c.JSON(resp.Code, resp)
	}

	data := structs.Map(result)

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}

// GetFindings for get list issue
func (adp *Adapter) CreateFindings(c echo.Context) error {
	ctx := context.Background()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	req := new(model.Findings)
	if err := c.Bind(req); err != nil {
		var resp Response
		errRes := adp.e.BadRequest(err.Error())
		resp.SetErrorResponse(errRes.Code, errRes.Message)
		return c.JSON(resp.Code, resp)
	}

	req.CreatedBy = claims["username"].(string)
	req.UpdatedBy = claims["username"].(string)

	act := findingsAct.NewCreateFindings(adp.repoPsql)
	result, err := act.Handle(ctx, *req)
	var resp Response

	if err != nil {
		resp.SetErrorResponse(err.Code, err.Message)
		return c.JSON(resp.Code, resp)
	}

	data := structs.Map(result)

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}
