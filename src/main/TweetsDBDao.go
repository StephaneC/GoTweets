package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"fmt"
	"encoding/json"
)

var (
	mgoSession   *mgo.Session
	databaseName = "tweets"
	collectionName = "tweets"
	databaseUri  = os.Getenv("MONGODB_ADDON_URI")
	err          error
)

func openDbConnection() *mgo.Session {
	if mgoSession == nil {
		var err error
		fmt.Printf("Open Connection on " + databaseUri)
		mgoSession, err = mgo.Dial(databaseUri)
		if err != nil {
			panic(err) // no, not really
			log.Fatal("FATAL : Error occured. Can't connect to database")
		}
	}
	return mgoSession.Clone()
}

/**
 * return tweets as JSON
 * 		or err as JSON too
 */
func getTweets(ts int64) string {
	fmt.Printf("Getting Tweets")
	c := mgoSession.DB(databaseName).C(collectionName)
	results := []Tweet{}
	aerr := c.Find(bson.M{"timestamp": bson.M{"$gt": ts}}).All(&results)
	if aerr != nil {
		log.Printf("Error getting tweets")
		return "{\"error\":\"Error occured\"}"
	}
	jsonTweets, _ := json.Marshal(results)
	fmt.Println(string(jsonTweets))
	return string(jsonTweets)
}

func addTweet(t string) {
	c := mgoSession.DB(databaseName).C(collectionName)
	err = c.Insert(t)
	if err != nil {
		log.Printf("Error")
	}
	log.Printf("tweet %s created\n", t)
}
