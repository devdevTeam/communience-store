package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func SearchFriend(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	colName := req.PostFormValue("colName")
	value := req.PostFormValue("value")
	friendList, err := lib.SearchFriend(uid, colName, value)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["friends"] = friendList
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
