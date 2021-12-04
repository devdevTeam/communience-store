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
	checkExist, err := CheckExistence(uid, rid)
	resp := make(map[string]interface{})
	if err != nil {
		ctx.Error(err)
		return
	}
	if checkExist {
		resp["error"] = "this user exists in this room"
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
		return
	}
	check, haveForm, err := CheckRoomPass(rid, password)
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
		if haveForm {
			err = lib.InsertNewCardValue(uid, rid)
			if err != nil {
				ctx.Error(err)
				return
			}
		}
		resp["error"] = err
		resp["haveForm"] = haveForm
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
	} else {
		resp["error"] = "isn't match password"
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
	}
}

func CheckExistence(uid, rid string) (bool, error) {
	relationInfo, err := lib.SelectUserRoomRelation(uid, rid)
	if err != nil {
		return false, err
	}
	if len(relationInfo) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func CheckRoomPass(rid, password string) (bool, bool, error) {
	roomInfo, err := lib.SelectRoom(rid)
	haveForm := roomInfo[3].(bool)
	if err != nil {
		return false, haveForm, err
	}

	if roomInfo[2] == password {
		return true, haveForm, nil
	} else {
		return false, haveForm, nil
	}
}
