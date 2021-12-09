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
	colName := req.PostFormValue("name")
	colHurigana := req.PostFormValue("hurigana")
	colBirthday := req.PostFormValue("birthday")
	colInstagram := req.PostFormValue("instagram")
	colTwitter := req.PostFormValue("twitter")
	colFacebook := req.PostFormValue("facebook")
	colFree := req.PostFormValue("free")
	colHobby := req.PostFormValue("hobby")
	err := lib.UpdateDefaultCard(uid, colName, colHurigana, colBirthday, colInstagram, colTwitter, colFacebook, colFree, colHobby)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
