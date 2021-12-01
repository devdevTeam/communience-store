package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func JoinEvent(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	eid := req.PostFormValue("eid")
	uid := req.PostFormValue("uid")
	password := req.PostFormValue("password")
	check, err := CheckEventPass(eid, password)
	if err != nil {
		ctx.Error(err)
		return
	}
	if check {
		err = lib.InsertParticipant(eid, uid)
		if err != nil {
			ctx.Error(err)
			return
		}
		resp := make(map[string]interface{})
		resp["error"] = err
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
	} else {
		resp := make(map[string]interface{})
		resp["error"] = "isn't match password"
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
	}
}

func CheckEventPass(eid, password string) (bool, error) {
	roomInfo, err := lib.SelectEvent(eid)
	if err != nil {
		return false, err
	}

	if roomInfo[2] == password {
		return true, nil
	} else {
		return false, nil
	}
}
