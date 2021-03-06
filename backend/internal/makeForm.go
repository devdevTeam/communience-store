package internal

import (
	"communience-store/backend/lib"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

func MakeForm(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	rid := req.PostFormValue("rid")
	colStr := req.PostFormValue("colList")
	colList := strings.Split(colStr, ",")
	err := lib.InsertNewForm(rid, colList)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
