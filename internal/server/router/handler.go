package router

import (
	"fmt"
	"net/http"
	"strconv"
	"zadanie/internal/app"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	app *app.App
}

func NewHandler(a *app.App) *Handler {
	return &Handler{
		app: a,
	}
}

// reserveGoods 	godoc
// @Summary			reserve goods
// @Descriprion		reserve array of goods by ID's, return reserve ID
// @Tags 			goods
// @Accept			json
// @Produce			json
// @Param			request body []int true "request body"
// @Success			200 	{string} 	string "reserve_id_example"
// @Failure			400 	{string} 	string "error"
// @Failure			500		{stirng}	string "error"
// @Router			/goods/reserve [post]
func (h *Handler) reserveGoods(ctx *gin.Context) {

	var reqBody []int64
	if err := ctx.BindJSON(&reqBody); err != nil {

		ctx.Error(fmt.Errorf("req body error"))
		ctx.String(http.StatusBadRequest, ctx.Errors.String())
		return
	}

	if len(reqBody) == 0 {

		ctx.Error(fmt.Errorf("empty goods id's array"))
		ctx.String(http.StatusBadRequest, ctx.Errors.String())
		return
	}

	reserveID, err := h.app.ReserveGoods(reqBody)
	if err != nil {
		ctx.Error(fmt.Errorf("reserve goods error"))
		ctx.String(http.StatusInternalServerError, ctx.Errors.String())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reserveID": reserveID})
}

// releaseGoods godoc
// @Summary			release goods
// @Descriprion 	release goods by ID's
// @Tags 			goods
// @Accept			json
// @Produce			json
// @Param			request body []int true "request body"
// @Success			200 	{string} 	string "ok"
// @Failure			400 	{string} 	string "error"
// @Failure			500		{stirng}	string "error"
// @Router			/goods/release [post]
func (h *Handler) releaseGoods(ctx *gin.Context) {

	var reqBody []int64
	if err := ctx.BindJSON(&reqBody); err != nil {

		ctx.Error(fmt.Errorf("req body error"))
		ctx.String(http.StatusBadRequest, ctx.Errors.String())
		return
	}

	if len(reqBody) == 0 {
		ctx.Error(fmt.Errorf("empty goods id's array"))
		ctx.String(http.StatusBadRequest, ctx.Errors.String())
		return
	}

	if err := h.app.ReleaseGoods(reqBody); err != nil {

		ctx.Error(fmt.Errorf("release goods error"))
		ctx.String(http.StatusInternalServerError, ctx.Errors.String())
		return
	}

	ctx.Status(http.StatusOK)

}

// getRemainGoods	godoc
// @Summary			get warehouse's goods
// @Descriprion 	get free warehouse's goods
// @Tags 			warehouse
// @Accept			json
// @Produce			json
// @Param			id 		path 	int	true	"warehouse_id" default(1)
// @Success			200 	{object} 	map[string]domain.Good	"{"goods":[{"id":3,"name":"mug","size":"6x6x7","quantity":15}]}"
// @Failure			400 	{string} 	string "error"
// @Failure			500		{stirng}	string "error"
// @Router			/warehouse/{id}/goods [get]
func (h *Handler) getRemainGoods(ctx *gin.Context) {
	id := ctx.Param("id")
	warehouseID, err := strconv.Atoi(id)
	if err != nil || warehouseID <= 0 {
		ctx.Error(fmt.Errorf("wrong warehouse id"))
		ctx.String(http.StatusBadRequest, ctx.Errors.String())
		return
	}

	goods, err := h.app.GetRemainGoods(int64(warehouseID))
	if err != nil {
		ctx.Error(fmt.Errorf("get remain goods error"))
		ctx.String(http.StatusInternalServerError, ctx.Errors.String())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"goods": goods})
}
