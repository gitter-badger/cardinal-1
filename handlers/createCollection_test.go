package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chasinglogic/cardinal/cards"
)

func TestCreateCollection(t *testing.T) {
	db := initTestDB()

	TestCard := cards.MagicCard{
		Name: "Test Card",
		Text: "This is a fake test card.",
	}

	mMC, err := json.Marshal(cards.MagicCollection{
		Name:   "testColl",
		Game:   "Magic",
		IsMain: false,
		Cards:  cards.MagicCard[TestCard],
	})

	req, _ := http.NewRequest("GET", "/api/v1/createCollection?user=test&game=magic&name=testColl", mMC)
	res := httptest.NewRecorder()

	CreateCollection(res, req, db)

	if res.Code != 200 {
		t.Fatalf("Expected: %v: Got: %v", "200", res.Code)
	}
}
