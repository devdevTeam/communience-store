package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetRoomList(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	roomList, err := lib.SelectUserRoomList(uid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["rooms"] = roomList
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
