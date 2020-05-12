package db

import (
	"config"
	scribble "github.com/nanobox-io/golang-scribble"
)

// top commodity
type TopCommodity string

// getDbClient returns a new scribble database, providing a destination for the database to live
func GetDbClient() *scribble.Driver{
	db, _ := scribble.New("./market", nil)
	return db
}

// updateDB will insert a record, the record will be an top Commodity.
func UpdateDB(db *scribble.Driver, name string ){
	t:= TopCommodity(name)
	db.Write(config.DBCollectionName, config.DBResourceName, t)
}

// getTopCommodityFromDB will read record from DB.
func GetTopCommodityFromDB(db *scribble.Driver) TopCommodity {
	t := TopCommodity("")
	db.Read(config.DBCollectionName,config.DBResourceName , &t)
	return t
}

