package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetRoomInfo(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	rid := req.PostFormValue("rid")
	roomInfo, err := lib.SelectRoom(rid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["rid"] = roomInfo[0]
	resp["name"] = roomInfo[1]
	resp["haveForm"] = roomInfo[3]
	resp["hash"] = roomInfo[4]
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
