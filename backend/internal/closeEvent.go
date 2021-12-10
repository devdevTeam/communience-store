package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func CloseEvent(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	eid := req.PostFormValue("eid")
	err := lib.DeleteEvents(eid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
