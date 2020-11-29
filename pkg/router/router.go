package router

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sulejman/scratchain/controller"
	"github.com/sulejman/scratchain/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(db *sql.DB) *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	api := controller.Controller{DB: db}

	// Non-protected routes
	posts := r.Group("/posts")
	{
		posts.GET("/address/:my_address", api.CalculateBalance)
		posts.POST("/transaction", api.CreateTransaction)
	}

	return r
}
