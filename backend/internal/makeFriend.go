package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func MakeFriend(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	uid := req.PostFormValue("uid")
	Shash := req.PostFormValue("hash")
	userInfo, err := lib.SelectUserWithHash(Shash)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	if len(userInfo) == 0 {
		resp["error"] = "invalid hash"
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
		return
	}
	checkExist, err := CheckExistenceFriend(uid, userInfo[0])
	if err != nil {
		ctx.Error(err)
		return
	}
	if checkExist {
		resp["error"] = "this user is your friend"
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
		return
	}
	err = lib.InsertUserFriendRelation(uid, userInfo[0])
	if err != nil {
		ctx.Error(err)
		return
	}
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}

func CheckExistenceFriend(uid, fid string) (bool, error) {
	relationInfo, err := lib.SelectUserFriendRelation(uid, fid)
	if err != nil {
		return false, err
	}
	if len(relationInfo) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
