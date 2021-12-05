package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetRoomUserDefaultCard(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	rid := req.PostFormValue("rid")
	defaultCardList, err := lib.SelectRoomUserDefaultCard(rid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["defaulCardList"] = defaultCardList
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
