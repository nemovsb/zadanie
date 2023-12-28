package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(h *Handler, mode string) (router *gin.Engine) {

	switch mode {
	case "dev":
		{
			gin.SetMode(gin.DebugMode)
		}
	case "prod":
		{
			gin.SetMode(gin.ReleaseMode)
		}
	}

	router = gin.New()
	router.Use(gin.Recovery())

	pprof.Register(router)

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	goods := router.Group("/goods")
	goods.POST("/reserve", h.reserveGoods)
	goods.POST("/release", h.releaseGoods)

	warehouse := router.Group("/warehouse")
	warehouse.GET(":id/goods", h.getRemainGoods)

	return router
}
