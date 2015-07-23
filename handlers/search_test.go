package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chasinglogic/cardinal/cards"
)

func TestCardSearch(t *testing.T) {
	db := initTestDB()

	req, _ := http.NewRequest("GET", "/api/v1/cardSearch?cardName=lightning%20bolt&game=magic", nil)
	res := httptest.NewRecorder()

	CardSearch(res, req, db)

	if res.Code != 200 {
		t.Fatalf("Expected: %v: Got: %v", "200", res.Code)
	}

	var c []cards.MagicCard

	err := json.Unmarshal(res.Body.Bytes(), &c)
	if err != nil {
		t.Fatalf("Unmarshal error: %s", err.Error())
	}

	for card := range c {
		t.Log(c[card])
	}
}
