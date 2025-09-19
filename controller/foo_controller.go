package controller

import (
	"net/http"
	"strconv"

	"github.com/Hidayathamir/goproj/pkg/http/response"
	"github.com/Hidayathamir/goproj/service"
	"github.com/gin-gonic/gin"
)

type FooController struct {
	fooService service.FooService
}

func NewFooController(fooService service.FooService) *FooController {
	return &FooController{
		fooService: fooService,
	}
}

// GetFoo godoc
//
//	@Summary		Get Foo by ID
//	@Description	Get a Foo record by its ID
//	@Tags			Foo
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"Foo ID"
//	@Success		200	{object}	response.ResHTTP[any]
//	@Failure		400	{object}	response.ResHTTP[any]
//	@Router			/foo [get]
func (c *FooController) GetFoo(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	req := service.GetFooReq{
		ID: id,
	}

	res, err := c.fooService.GetFoo(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Data(ctx, http.StatusOK, res)
}

// SaveFoo godoc
//
//	@Summary		Save Foo data
//	@Description	Save a new Foo record
//	@Tags			Foo
//	@Accept			json
//	@Produce		json
//	@Param			data	query		string	true	"Data to save"
//	@Success		200		{object}	response.ResHTTP[any]
//	@Failure		400		{object}	response.ResHTTP[any]
//	@Router			/foo [post]
func (c *FooController) SaveFoo(ctx *gin.Context) {
	data := ctx.Query("data")

	req := service.SaveFooReq{
		Data: data,
	}

	res, err := c.fooService.SaveFoo(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Data(ctx, http.StatusOK, res)
}

// BackupFoo godoc
//
//	@Summary		Backup Foo by ID
//	@Description	Backup a Foo record
//	@Tags			Foo
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"Foo ID"
//	@Success		200	{object}	response.ResHTTP[any]
//	@Failure		400	{object}	response.ResHTTP[any]
//	@Router			/foo/backup [post]
func (c *FooController) BackupFoo(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	req := service.BackupFooReq{
		ID: id,
	}

	res, err := c.fooService.BackupFoo(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Data(ctx, http.StatusOK, res)
}

// PayForFoo godoc
//
//	@Summary		Pay for Foo
//	@Description	Perform a payment for a Foo record
//	@Tags			Foo
//	@Accept			json
//	@Produce		json
//	@Param			id			query		int		true	"Foo ID"
//	@Param			user		query		string	true	"Username"
//	@Param			pass		query		string	true	"Password"
//	@Param			amount		query		number	true	"Amount"
//	@Param			currency	query		string	true	"Currency"
//	@Success		200			{object}	response.ResHTTP[any]
//	@Failure		400			{object}	response.ResHTTP[any]
//	@Router			/foo/pay [post]
func (c *FooController) PayForFoo(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	username := ctx.Query("user")
	password := ctx.Query("pass")
	amountStr := ctx.Query("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	currency := ctx.Query("currency")

	req := service.PayForFooReq{
		ID:       id,
		Username: username,
		Password: password,
		Amount:   amount,
		Currency: currency,
	}

	res, err := c.fooService.PayForFoo(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Data(ctx, http.StatusOK, res)
}
