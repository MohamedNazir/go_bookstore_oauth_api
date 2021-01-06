package app

import (
	"github.com/MohamedNazir/go_bookstore_oauth_api/src/domain/access_token"
	"github.com/MohamedNazir/go_bookstore_oauth_api/src/http"
	"github.com/MohamedNazir/go_bookstore_oauth_api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atservice := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atservice)

	router.GET("/oauth/accesstoken/:tokenid", atHandler.GetByID)
	router.POST("/oauth/accesstoken", atHandler.Create)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router.Run(":8080")
}
