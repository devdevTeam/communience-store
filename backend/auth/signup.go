package auth

import (
	"encoding/json"
	"log"

	firebaselib "communience-store/backend/firebase"
	"communience-store/backend/lib"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// Signup create User
func Signup(ctx *gin.Context) {
	req := ctx.Request
	req.ParseForm()
	mail := req.PostFormValue("mail")
	password := req.PostFormValue("password")
	userName := req.PostFormValue("username")
	err := signupValidator(mail, password, userName)
	if err != nil {
		ctx.Error(err)
		return
	}
	firebaselib.Ini()
	err = createUser(mail, password, userName)
	resp := make(map[string]interface{})
	resp["error"] = err
	res, _ := json.Marshal(resp)
	ctx.Writer.Write(res)
}

func createUser(mail, password, username string) error {
	params := (&auth.UserToCreate{}).
		Email(mail).
		EmailVerified(firebaselib.Verified).
		Password(password).
		DisplayName(username).
		Disabled(firebaselib.Disabled)
	u, err := firebaselib.Client.CreateUser(firebaselib.Ctx, params)
	if err != nil {
		return err
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)
	lib.InsertNewUser(u.UserInfo.UID, u.UserInfo.Email, password, u.UserInfo.DisplayName)
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
