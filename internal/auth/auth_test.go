package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	id := uuid.New()
	_, err := MakeJWT(id, "test", 5000)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateJWT(t *testing.T) {
	id := uuid.New()
	token, err := MakeJWT(id, "test", time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	retId, err := ValidateJWT(token, "test")
	if err != nil {
		t.Fatal(err)
	}
	if retId != id {
		t.Fatal("id should be equal")
	}
}

func TestGetBearerToken(t *testing.T) {
	token, err := GetBearerToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	retId, err := ValidateJWT(token, "test")
	if err != nil {
		t.Fatal(err)
	}
	if retId == uuid.Nil {
		t.Fatal("token should not be nil")
	}
}

func TestMakeRefreshToken(t *testing.T) {
}
