package nativestore

import (
	"log"
	"testing"
)

func TestSetGet(t *testing.T) {
	return
	url := "https://github.com/mobingi/gosdk"
	Set("gosdk", url, "user", "password")
	user, secret, err := Get("gosdk", url)
	if err == nil {
		if user != "user" {
			t.Errorf("Expecting user, got %s", user)
		}

		if secret != "password" {
			t.Errorf("Expecting password, got %s", secret)
		}
	} else {
		log.Println("got error:", err)
	}
}
