package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetRoomUsers(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("rid")
	userList, err := lib.SelectRoomUsers(uid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["users"] = userList
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
