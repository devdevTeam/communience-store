package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var Ctx = context.Background()
var opt = option.WithCredentialsFile("/app/backend/firebase/key/communience-store-key.json")
var App *firebase.App
var Client *auth.Client

const Verified = false
const Disabled = false

func Ini() {
	var err error
	App, err = firebase.NewApp(context.Background(), nil, opt)
	must(err)
	Client, err = App.Auth(Ctx)
	must(err)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
