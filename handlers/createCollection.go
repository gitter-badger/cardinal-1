package handlers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"

	logger "github.com/Sirupsen/logrus"
	"github.com/chasinglogic/cardinal/cards"
)

// CreateCollection will add a collection to the appropriate user
func CreateCollection(w http.ResponseWriter, r *http.Request, db *mgo.Database) {
	logger.Info("create collection handler called.")
	decoder := json.NewDecoder(r.Body)
	logger.Info("past decoder")
	var mc cards.MagicCollection
	logger.Info("mc made")

	err := decoder.Decode(&mc)
	logger.Info("decoded")
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info(mc)
}
