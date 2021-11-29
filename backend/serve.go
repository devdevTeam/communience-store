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
	api.POST("/getRoomList", internal.GetRoomList)
	api.POST("/makeForm", internal.MakeForm)
	api.POST("/getForm", internal.GetForm)
	api.POST("/startEvent", internal.StartEvent)
	api.POST("/getEventList", internal.GetEventList)

	e := r.Group("/api/event")
	internal.NewMelody()
	e.POST("/joinEvent", internal.JoinEvent)
	e.GET("/:eid", internal.EventWebsocket)
	internal.DefineMelodyBehavior()

	r.GET("/", root)
	port := ":8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	r.Run(":" + port)
}

func root(ctx *gin.Context) {
	ctx.Writer.WriteString("hello")
}
