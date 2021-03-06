package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetForm(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	rid := req.PostFormValue("rid")
	colList, err := lib.SelectForm(rid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["colList"] = colList
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
