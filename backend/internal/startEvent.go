package internal

import (
	"communience-store/backend/lib"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func StartEvent(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	password := req.PostFormValue("password")
	org_uid := req.PostFormValue("org_uid")
	rid := req.PostFormValue("rid")
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		ctx.Error(err)
		return
	}
	eid := u.String()
	err = lib.InsertNewEvent(eid, password, org_uid, rid)
	if err != nil {
		ctx.Error(err)
		return
	}
	colList, err := lib.SelectForm(rid)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = lib.InsertEventCol(eid, colList)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}
