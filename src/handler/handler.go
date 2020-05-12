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
	// swagger:operation GET /api/v1/GetCommodity GetCommodity
	//
	// Returns Top Commodity item from stock market.
	// It will cache the item until data is updated
	// in DB, Then receiver will get notification
	// which will update the cache itself.
	// ---
	// consumes:
	// - text/plain
	// produces:
	// - text/plain
	// responses:
	//   '200':
	//     description: Top Commodity Item from the stock market
	//     type: string
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
	// swagger:operation POST /api/v1/InitializeCommodity InitializeCommodity
	//
	// Fill the cache with data from DB.
	// ---
	// consumes:
	// - text/plain
	// produces:
	// - text/plain
	// responses:
	//   '200':
	//     description: update the cache with data from DB.
	//     type: string
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