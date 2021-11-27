package internal

import (
	"communience-store/backend/lib"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MakeRoom(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	roomName := req.PostFormValue("roomname")
	password := req.PostFormValue("password")
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		ctx.Error(err)
		return
	}
	rid := u.String()
	err = lib.InsertNewRoom(rid, roomName, password)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = lib.InsertUserRoomRelation(uid, rid, true)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
