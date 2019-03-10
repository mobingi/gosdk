package session

import (
	"log"
	"os"
	"testing"
)

func TestNewSession(t *testing.T) {
	s1, _ := New()
	if s1 == nil {
		t.Errorf("Expected non-nil session")
	}

	s2, _ := New(&Config{})
	if s2 == nil {
		t.Errorf("Expected non-nil session")
	}

	s3, _ := New(&Config{
		ClientId:     "clientid",
		ClientSecret: "clientsecret",
	})

	if s3.Config.ClientId != "clientid" {
		t.Errorf("Expected value 'clientid', got %s", s3.Config.ClientId)
	}

	s4, _ := New(&Config{ApiVersion: 3})
	if s4.ApiEndpoint() != "https://api.mobingi.com/v3" {
		t.Errorf("Invalid api url")
	}
}

func TestNewSessionDevAcct(t *testing.T) {
	return
	if os.Getenv("MOBINGI_CLIENT_ID") != "" && os.Getenv("MOBINGI_CLIENT_SECRET") != "" {
		s, err := New(&Config{
			BaseApiUrl: "https://logindev.mobingi.com",
			GrantType:  "client_credentials",
			ApiVersion: -1, // no version
		})

		if err != nil {
			t.Errorf("Should succeed, got %v", err)
		}

		if s.AccessToken == "" {
			t.Errorf("Should have token, got empty")
		}

		log.Println(s)
	}
}

func TestNewSessionDevAcctOld(t *testing.T) {
	return
	if os.Getenv("MOBINGI_CLIENT_ID") != "" && os.Getenv("MOBINGI_CLIENT_SECRET") != "" {
		s, err := New(&Config{
			BaseApiUrl: "https://apidev.mobingi.com",
		})

		if err != nil {
			t.Errorf("Should succeed, got %v", err)
		}

		if s.AccessToken == "" {
			t.Errorf("Should have token, got empty")
		}

		log.Println(s)
	}
}
