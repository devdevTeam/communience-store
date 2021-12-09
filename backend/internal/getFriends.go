package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetUserFriends(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	friendList, err := lib.SelectUserFriend(uid)
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
