package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func JoinRoom(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	rid := req.PostFormValue("rid")
	password := req.PostFormValue("password")
	check, err := CheckRoomPass(rid, password)
	if err != nil {
		ctx.Error(err)
		return
	}
	if check {
		err = lib.InsertUserRoomRelation(uid, rid, false)
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

func CheckRoomPass(rid, password string) (bool, error) {
	roomInfo, err := lib.SelectRoom(rid)
	if err != nil {
		return false, err
	}

	if roomInfo[2] == password {
		return true, nil
	} else {
		return false, nil
	}
}
