package handlers

import (
	"net/http"

	"gopkg.in/mgo.v2"

	logger "github.com/Sirupsen/logrus"
)

// CreateCollection will add a collection to the appropriate user
func CreateCollection(w http.ResponseWriter, r *http.Request, db *mgo.Database) {
	logger.Debug("create collection handler called.")
	//decoder := json.NewDecoder(r.Body)
	//var cc cards.Collection

	//err := decoder.Decode(&cc)
}
