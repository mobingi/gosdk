package session

import (
	"log"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
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
	if os.Getenv("MOBINGI_OPENID_CLIENT_ID") != "" && os.Getenv("MOBINGI_OPENID_CLIENT_SECRET") != "" {
		s, err := New(&Config{
			ClientId:     os.Getenv("MOBINGI_OPENID_CLIENT_ID"),
			ClientSecret: os.Getenv("MOBINGI_OPENID_CLIENT_SECRET"),
			BaseApiUrl:   "https://logindev.mobingi.com",
			GrantType:    "client_credentials",
			ApiVersion:   -1, // no version
			UseForm:      true,
			Scope:        "openid",
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

func testTokenFromSession(t *testing.T, s *Session) {
	if s == nil {
		t.Fatal("should not be nil")
	}

	if s.AccessToken == "" {
		t.Fatal("access token empty")
	}

	log.Println("token:", s.AccessToken)
}

func TestNewSession(t *testing.T) {
	return
	s1 := NewSession(
		WithClientId(os.Getenv("MOBINGI_OPENID_CLIENT_ID")),
		WithClientSecret(os.Getenv("MOBINGI_OPENID_CLIENT_SECRET")),
		WithBaseLoginUrl("https://logindev.mobingi.com"),
	)

	testTokenFromSession(t, s1)

	s2 := NewSession(
		WithClientId(os.Getenv("MOBINGI_OPENID_CLIENT_ID")),
		WithClientSecret(os.Getenv("MOBINGI_OPENID_CLIENT_SECRET")),
		WithBaseLoginUrl("https://logindev.mobingi.com"),
		WithGrantType("password"),
		WithUsername(os.Getenv("MOBINGI_USERNAME")),
		WithPassword(os.Getenv("MOBINGI_PASSWORD")),
	)

	testTokenFromSession(t, s2)
}
