package controller

import (
	"net/http"

	"github.com/Hidayathamir/goproj/pkg/http/response"
	"github.com/Hidayathamir/goproj/service"
	"github.com/gin-gonic/gin"
)

type BarController struct {
	barService service.BarService
}

func NewBarController(barService service.BarService) *BarController {
	return &BarController{
		barService: barService,
	}
}

// DoSomething godoc
//
//	@Summary		Perform something with user credentials
//	@Description	Calls DoSomething service with user and password
//	@Tags			Bar
//	@Accept			json
//	@Produce		json
//	@Param			user	query		string	true	"Username"
//	@Param			pass	query		string	true	"Password"
//	@Success		200		{object}	response.ResHTTP[any]
//	@Router			/bar/do-something [get]
func (c *BarController) DoSomething(ctx *gin.Context) {
	user := ctx.Query("user")
	pass := ctx.Query("pass")

	req := service.DoSomethingReq{
		User: user,
		Pass: pass,
	}

	res, err := c.barService.DoSomething(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Data(ctx, http.StatusOK, res)
}

// DoAnotherThing godoc
//
//	@Summary		Perform another action using a token
//	@Description	Calls DoAnotherThing service with a token
//	@Tags			Bar
//	@Accept			json
//	@Produce		json
//	@Param			token	query		string	true	"Access token"
//	@Success		200		{object}	response.ResHTTP[any]
//	@Router			/bar/do-another [get]
func (c *BarController) DoAnotherThing(ctx *gin.Context) {
	token := ctx.Query("token")

	req := service.DoAnotherThingReq{
		Token: token,
	}

	res, err := c.barService.DoAnotherThing(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Data(ctx, http.StatusOK, res)
}

// GetAllData godoc
//
//	@Summary		Get all data using a token
//	@Description	Returns all data from the service for a given token
//	@Tags			Bar
//	@Accept			json
//	@Produce		json
//	@Param			token	query		string	true	"Access token"
//	@Success		200		{object}	response.ResHTTP[any]
//	@Router			/bar/all-data [get]
func (c *BarController) GetAllData(ctx *gin.Context) {
	token := ctx.Query("token")

	req := service.GetAllDataReq{
		Token: token,
	}

	res, err := c.barService.GetAllData(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Data(ctx, http.StatusOK, res)
}

// NotifyAll godoc
//
//	@Summary		Send notification to all users
//	@Description	Sends a message using a token; returns error if sending fails
//	@Tags			Bar
//	@Accept			json
//	@Produce		json
//	@Param			token	query		string	true	"Access token"
//	@Param			message	query		string	true	"Message content"
//	@Success		200		{object}	response.ResHTTP[any]
//	@Failure		500		{object}	response.ResHTTP[any]
//	@Router			/bar/notify [post]
func (c *BarController) NotifyAll(ctx *gin.Context) {
	token := ctx.Query("token")
	message := ctx.Query("message")

	req := service.NotifyAllReq{
		Token:   token,
		Message: message,
	}

	res, err := c.barService.NotifyAll(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Data(ctx, http.StatusOK, res)
}
