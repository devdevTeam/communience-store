package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func JoinRoom(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	token := req.PostFormValue("token")
	rid := req.PostFormValue("rid")
	uid := lib.ConvertUID(token)
	err := lib.InsertUserRoomRelation(uid, rid, false)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
