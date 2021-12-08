package internal

import (
	"communience-store/backend/lib"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchRoomUser(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	rid := req.PostFormValue("rid")
	haveForm := req.PostFormValue("haveForm")
	BhaveForm, _ := strconv.ParseBool(haveForm)
	colName := req.PostFormValue("colName")
	value := req.PostFormValue("value")
	var userList []interface{}
	var err error
	if BhaveForm {
		userList, err = lib.SearchRoomUser_form(rid, colName, value)
	} else {
		userList, err = lib.SearchRoomUser_default(rid, colName, value)
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
