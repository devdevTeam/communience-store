package internal

import (
	"communience-store/backend/lib"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func JoinEvent(eid, uid, password string) (interface{}, error) {
	check, err := CheckRoomPass(eid, password)
	if err != nil {
		return nil, err
	}
	if check {
		err = lib.InsertParticipant(eid, uid)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else {
		resp := make(map[string]interface{})
		resp["error"] = "isn't match password"
		return resp, nil
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

func FetchEventCol(eid string) (interface{}, error) {
	eventCol, err := lib.SelectEventCol(eid)
	if err != nil {
		return nil, err
	}
	resp := make(map[string]interface{})
	resp["error"] = err
	resp["colList"] = eventCol
	return resp, nil
}

var m *melody.Melody

func NewMelody() {
	m = melody.New()
}

func EventWebsocket(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	eid := req.PostFormValue("eid")
	uid := req.PostFormValue("uid")
	password := req.PostFormValue("password")
	resp, err := JoinEvent(eid, uid, password)
	if resp != nil {
		res, _ := json.Marshal(resp)
		ctx.Writer.Write(res)
		return
	}
	if err != nil {
		ctx.Error(err)
		return
	}
	m.HandleRequest(ctx.Writer, ctx.Request)
}

type Msg struct {
	Eid    string `json:"eid"`
	UID    string `json:"uid"`
	Hidden []bool `json:"hidden"`
}

func eventHandler(s1 *melody.Session, msg []byte) {
	var messageObj Msg
	json.Unmarshal(msg, &messageObj)
	resp, err := FetchEventCol(messageObj.Eid)
	if err != nil {
		s1.Write([]byte(err.Error()))
	}
	res, _ := json.Marshal(resp)
	m.BroadcastFilter(res, func(s2 *melody.Session) bool {
		return s1.Request.URL.Path == s2.Request.URL.Path
	})
}

func DefineMelodyBehavior() {
	m.HandleMessage(eventHandler)
}
