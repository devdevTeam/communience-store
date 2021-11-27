package main

import (
	"communience-store/backend/auth"
	"communience-store/backend/internal"
	"communience-store/backend/lib"
	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	lib.DBInit()
	r := gin.Default()

	// cors
	config := cors.DefaultConfig()
	r.Use(cors.New(config))

	// routing
	api := r.Group("/api")
	api.POST("/signup", auth.Signup)
	api.POST("/makeRoom", internal.MakeRoom)
	api.POST("/joinRoom", internal.JoinRoom)
	api.POST("/leaveRoom", internal.LeaveRoom)
	api.POST("/getUserInfo", internal.GetUserInfo)

	r.GET("/", root)
	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	r.Run(":" + port)
}

func root(ctx *gin.Context) {
	ctx.Writer.WriteString("hello")
}
