package memory

import (
	"context"
	"testing"

	"github.com/webeauford/accounts/pkg/accounts"
)

func TestGetUser(t *testing.T) {
	svc := New()
	svc.PutUser(accounts.User{ID: "1", Email: "a@b.com", Name: "Will"})

	u, err := svc.GetUser(context.Background(), "1")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if u.Email != "a@b.com" {
		t.Fatalf("expected email a@b.com, got %s", u.Email)
	}
}

func TestGetUser_NotFound(t *testing.T) {
	svc := New()

	_, err := svc.GetUser(context.Background(), "nope")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
