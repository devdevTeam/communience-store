package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func CheckHash(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	Shash := req.PostFormValue("hash")
	roomInfo, err := lib.SelectRoomWithHash(Shash)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	if len(roomInfo) == 0 {
		resp["error"] = "invalid hash"
	} else {
		resp["error"] = err
	}
	resp["roomInfo"] = roomInfo
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
