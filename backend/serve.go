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
	api.POST("/getUserInfo", internal.GetUserInfo)
	api.POST("/getDefaultCard", internal.GetDefaultCard)
	api.POST("/updateDefaultCard", internal.UpdateDefaultCard)
	api.POST("/makeRoom", internal.MakeRoom)
	api.POST("/getRoomInfo", internal.GetRoomInfo)
	api.POST("/getRoomList", internal.GetRoomList)
	api.POST("/searchRoom", internal.SearchRoom)
	api.POST("/getRoomUsers", internal.GetRoomUsers)
	api.POST("/searchRoomUser", internal.SearchRoomUser)
	api.POST("/getRoomAdmin", internal.GetRoomAdmin)
	api.POST("/getRoomUserDefaultCard", internal.GetRoomUserDefaultCard)
	api.POST("/getCardValue", internal.GetCardValue)
	api.POST("/updateCardValue", internal.UpdateCardValue)
	api.POST("/joinRoom", internal.JoinRoom)
	api.POST("/checkHash", internal.CheckHash)
	api.POST("/leaveRoom", internal.LeaveRoom)
	api.POST("/makeForm", internal.MakeForm)
	api.POST("/getForm", internal.GetForm)
	api.POST("/startEvent", internal.StartEvent)
	api.POST("/closeEvent", internal.CloseEvent)
	api.POST("/getEventList", internal.GetEventList)
	api.POST("/makeFriend", internal.MakeFriend)
	api.POST("/getFriends", internal.GetUserFriends)
	api.POST("/searchFriend", internal.SearchFriend)

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
