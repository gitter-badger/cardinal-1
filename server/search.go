package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

type results struct {
	Docs []MagicCard
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchTerm := vars["name"]
	game := vars["game"]
	var result []MagicCard
	logger.Debug(game)
	logger.Debug(searchTerm)

	if game == "magic" {

		ferr := magicCollection.Find(bson.M{"name": &bson.M{"$regex": ".*" + searchTerm + ".*", "$options": "i"}}).All(&result)
		if ferr != nil {
			logger.Warn("Find Error")
			errCheck(ferr)
		}

		logger.Debug(result)

		marshaledResults, merr := json.Marshal(result)
		errCheck(merr)

		w.Write(marshaledResults)
	} else if game == "hearthstone" {
		// Not quite ready for hearthstone
		w.WriteHeader(500)
	}
}
