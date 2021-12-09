package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetDefaultCard(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	cardInfo, err := lib.SelectDefaultCard(uid)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["name"] = cardInfo[0]
	resp["hurigana"] = cardInfo[1]
	resp["birthday"] = cardInfo[2]
	resp["instagram"] = cardInfo[3]
	resp["twitter"] = cardInfo[4]
	resp["facebook"] = cardInfo[5]
	resp["free"] = cardInfo[6]
	resp["hobby"] = cardInfo[7]
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
