package internal

import (
	"communience-store/backend/lib"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

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

func connected(s1 *melody.Session) {
	fmt.Println("connected! : ", s1.Request.URL.Path)
}

func DefineMelodyBehavior() {
	m.HandleConnect(connected)
	m.HandleMessage(eventHandler)
}
