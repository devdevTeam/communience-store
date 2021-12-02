package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetCardValue(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	rid := req.PostFormValue("rid")
	cardValue, err := lib.SelectCardValue(uid, rid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["cardValue"] = cardValue
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
