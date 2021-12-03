package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func SearchRoom(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	name := req.PostFormValue("name")
	rid := req.PostFormValue("rid")
	roomList, err := lib.SearchRoom(name, rid)
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
