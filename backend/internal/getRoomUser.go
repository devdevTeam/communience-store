package internal

import (
	"communience-store/backend/lib"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRoomUsers(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	rid := req.PostFormValue("rid")
	haveForm := req.PostFormValue("haveForm")
	BhaveForm, _ := strconv.ParseBool(haveForm)
	var userList []interface{}
	var err error
	if BhaveForm {
		userList, err = lib.SelectRoomDisplayInfo_forms(rid)
	} else {
		userList, err = lib.SelectRoomDisplayInfo_default(rid)
	}
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
