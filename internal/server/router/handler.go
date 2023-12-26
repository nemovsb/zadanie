package router

import (
	"fmt"
	"net/http"
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

func (h *Handler) reserveGoods(ctx *gin.Context) {

	var reqBody []string
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(fmt.Errorf("req body error"))
		return
	}

	if err := h.app.ReserveGoods(reqBody); err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Error(fmt.Errorf("reserve goods error"))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) releaseGoods(ctx *gin.Context) {

	var reqBody []string
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(fmt.Errorf("req body error"))
		return
	}

	if err := h.app.ReleaseGoods(reqBody); err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Error(fmt.Errorf("release goods error"))
		return
	}

	ctx.Status(http.StatusOK)

}

func (h *Handler) getRemainGoods(ctx *gin.Context) {
	warehouseID := ctx.Param("id")
	if warehouseID == "" {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(fmt.Errorf("empty warehouse id"))
		return
	}

	goods, err := h.app.GetRemainGoods(warehouseID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Error(fmt.Errorf("get remain goods error"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"goods": goods})
}
