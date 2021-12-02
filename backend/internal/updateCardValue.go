package internal

import (
	"communience-store/backend/lib"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateCardValue(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	rid := req.PostFormValue("rid")
	valueStr := req.PostFormValue("valueList")
	valueList := strings.Split(valueStr, ",")
	err := lib.UpdateCardValue(uid, rid, valueList)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
