package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	userInfo, err := lib.SelectUser(uid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["uid"] = userInfo[0]
	resp["mail"] = userInfo[1]
	resp["name"] = userInfo[3]
	resp["hash"] = userInfo[4]
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
