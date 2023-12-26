package router

import (
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

}

func (h *Handler) releaseGoods(ctx *gin.Context) {

}

func (h *Handler) getRemainGoods(ctx *gin.Context) {

}
