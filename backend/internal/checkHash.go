package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func CheckHash(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	Shash := req.PostFormValue("hash")
	roomInfo, err := lib.SelectRoomWithHash(Shash)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	if len(roomInfo) == 0 {
		resp["error"] = "invalid hash"
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
		return
	}
	msg, haveForm, err := join(uid, roomInfo[0], roomInfo[1])
	if err != nil {
		ctx.Error(err)
		return
	}
	resp["error"] = msg
	resp["rid"] = roomInfo[0]
	resp["haveForm"] = haveForm
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}

func join(uid, rid, password string) (string, bool, error) {
	check, haveForm, err := check(rid, password)
	if err != nil {
		return "", false, err
	}
	if check {
		err = lib.InsertUserRoomRelation(uid, rid, false)
		if err != nil {
			return "", false, err
		}
		if haveForm {
			err = lib.InsertNewCardValue(uid, rid)
			if err != nil {
				return "", false, err
			}
		}
		return "", haveForm, nil
	} else {
		return "isn't match password", false, err
	}
}

func check(rid, password string) (bool, bool, error) {
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
