package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func UpdateDefaultCard(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	colName := req.PostFormValue("colName")
	colHurigana := req.PostFormValue("colHurigana")
	colBirthday := req.PostFormValue("colBirthday")
	colInstagram := req.PostFormValue("colInstagram")
	colTwitter := req.PostFormValue("colTwitter")
	colFacebook := req.PostFormValue("colFacebook")
	colFree := req.PostFormValue("colFree")
	err := lib.UpdateDefaultCard(uid, colName, colHurigana, colBirthday, colInstagram, colTwitter, colFacebook, colFree)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
