package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chasinglogic/cardinal/cards"
)

func TestCreateCollection(t *testing.T) {
	db := initTestDB()

	mMC, err := json.Marshal(cards.MagicCollection{
		Name:   "testColl",
		Game:   "Magic",
		IsMain: false,
		Cards: []cards.MagicCard{cards.MagicCard{
			Name: "Test Card",
			Text: "This is a fake test card.",
		}},
	})
	if err != nil {
		t.Fatal("Marshal Error: " + err.Error())
	}

	req, _ := http.NewRequest("GET", "/api/v1/createCollection?user=test", bytes.NewReader(mMC))
	res := httptest.NewRecorder()

	CreateCollection(res, req, db.C("users"))

	if res.Code != 200 {
		t.Fatalf("Expected: %v: Got: %v", "200", res.Code)
	}
}
