package handler

import (
	"config"
	"db"
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"net/http"
)

var Cache = cache.New(cache.NoExpiration,cache.NoExpiration)

func GetCommodity(w http.ResponseWriter, _ *http.Request) {

	data, found := Cache.Get(config.DBCollectionName)
	if found {
		  d  := data.(db.TopCommodity)
		bytes, err := json.Marshal( d )
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		writeJsonResponse(w, bytes)
	}

}

func InitilizeCommodity(w http.ResponseWriter, _ *http.Request) {
	UpdateCache()
	w.WriteHeader(http.StatusCreated)
}

func UpdateCache(){
	dbClient := db.GetDbClient()
	Cache.Set(config.DBCollectionName,db.GetTopCommodityFromDB(dbClient),cache.NoExpiration)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}