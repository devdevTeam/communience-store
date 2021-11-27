package lib

import (
	firebaselib "communience-store/backend/firebase"
)

func ConvertUID(token string) string {
	firebaselib.Ini()
	client, _ := firebaselib.App.Auth(firebaselib.Ctx)
	tkn, _ := client.VerifyIDToken(firebaselib.Ctx, token)
	return tkn.UID
}
