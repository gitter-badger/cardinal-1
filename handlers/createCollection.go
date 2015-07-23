package handlers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	logger "github.com/Sirupsen/logrus"
	"github.com/chasinglogic/cardinal/cards"
)

// CreateCollection will add a collection to the appropriate user
func CreateCollection(w http.ResponseWriter, r *http.Request, collection *mgo.Collection) {
	logger.Info("Creating collection.")
	user := r.FormValue("user")

	decoder := json.NewDecoder(r.Body)
	var mc cards.MagicCollection

	err := decoder.Decode(&mc)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(500)
	}

	uerr := collection.Update(bson.M{"username": user}, bson.M{
		"$push": bson.M{
			"collections": mc,
		},
	})
	if uerr != nil {
		logger.Error(err.Error())
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}
