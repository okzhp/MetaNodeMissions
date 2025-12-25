package router

import (
	"auctionBackEnd/controller"
	"auctionBackEnd/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	ctl := controller.AuctionController{Db: db}
	v1 := router.Group("/api/v1")
	{
		v1.GET("/auction", ctl.QueryAuction)
		v1.GET("/bid", ctl.QueryBid)
		v1.GET("/health", func(c *gin.Context) {
			utils.Success(c, "api is running")
		})
	}

	return router

}
