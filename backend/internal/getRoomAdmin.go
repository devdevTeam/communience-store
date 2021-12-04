package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetRoomAdmin(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	rid := req.PostFormValue("rid")
	relationInfo, err := lib.SelectUserRoomRelation(uid, rid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["admin"] = relationInfo["admin"]
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
