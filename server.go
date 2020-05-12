package main

import (
	"config"
	"db"
	"github.com/gorilla/mux"
	"handler"
	"kafka"
	"log"
	"net/http"
	"time"
)


func main(){


	go InitilizeAndUpdateDB()
	go kafka.WaitForNotification()
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix(config.PathPrefix).Subrouter()
	sub.Methods("GET").Path(config.GetMethodPath).HandlerFunc(handler.GetCommodity)
	sub.Methods("POST").Path(config.InitializeMethodPath).HandlerFunc(handler.InitilizeCommodity)

	log.Fatal(http.ListenAndServe(":3000", router))
}


// InitilizeAndUpdateDB will Initialize the DB with some values and update it the record after sleeping every 10 seconds.
func InitilizeAndUpdateDB(){
	for{

		dbClient := db.GetDbClient()
		//update the top commodity to database.
		for _, name := range config.CommudityData {
			db.UpdateDB(dbClient,name)
			time.Sleep(time.Duration(config.SleepTime) * time.Second)
			kafka.NotifyReceivers()
		}
	}
}

