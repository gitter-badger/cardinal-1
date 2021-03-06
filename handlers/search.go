package handlers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	logger "github.com/Sirupsen/logrus"
	"github.com/chasinglogic/cardinal/cards"
)

// CardSearch is an HTTP Handler which searches for a card. Takes the game and name of the card from the url
func CardSearch(w http.ResponseWriter, r *http.Request, db *mgo.Database) {
	searchTerm := r.FormValue("cardName")
	game := r.FormValue("game")
	logger.Debug("Searching for " + searchTerm + " in " + game)

	if game == "hearthstone" {
		// Not quite ready for hearthstone
		w.WriteHeader(http.StatusNotImplemented)
	}

	var result []cards.MagicCard
	ferr := db.C(game).Find(bson.M{"name": &bson.M{"$regex": ".*" + searchTerm + ".*", "$options": "i"}}).All(&result)
	if ferr != nil {
		logger.Error(ferr)
		w.WriteHeader(http.StatusNotFound)
	}

	marshaledResults, merr := json.Marshal(result)
	if merr != nil {
		logger.Error(ferr)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(marshaledResults)
}
