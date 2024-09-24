package utils

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type FlashMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func SessionFlash(store *session.Store, ctx *fiber.Ctx, message fiber.Map) {
	sess, err := store.Get(ctx)
	if err != nil {
		panic(err)
	}

	flashMessage := FlashMessage{
		Message: message["message"].(string),
		Type:    message["type"].(string),
	}

	b, err2 := json.Marshal(flashMessage)
	if err2 != nil {
		panic(err2)
	}

	sess.Set("flash", b)
	err = sess.Save()
	if err != nil {
		panic(err)
	}
}
