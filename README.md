# A sample Cache Go microservice

This is a simple Microservice which will provide top commodity (Iron, Gold, Platinium and silver).
Gorilla is used as web toolkit, It makes use of go-cache to store the data in cache-memory,
scribble is used as database, Any update in the database Apache Kafaka will notify so that cache will be updated automatically.
 

## Packages

* gorilla : A Web Toolkit.
* Sarama : Go client library for Apache Kafka.
* scribble : A tiny JSON database in Golang.
* go-cache : An in-memory key:value store.

## Usage 

* Clone the Repo
```
git clone https://github.com/akshayvijapur/Go_sample_microservice.git
```
* Go to the repo directory
```
cd Go_sample_microservice/
```
* perform docker compose up
``` 
docker-compose up -d
```
* Initialize the Cache by calling to InitializeCommand REST API

```
curl -X POST http://localhost:3000/api/v1/InitializeCommodity
``` 

* Now Cache will be updated with Commodity data from DB.

* Get the Latest Commodity by calling GetCommodity  REST API

```
curl http://localhost:3000/api/v1/GetCommodity
```

* As DB is being updated with new Commodity, Kafka will notify the receiver 
which in turns will update the cache. 