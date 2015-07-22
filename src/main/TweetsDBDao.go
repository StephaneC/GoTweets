package main
/*
import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"os"
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
		mgoSession, err = mgo.Dial(databaseUri)
		if err != nil {
			panic(err) // no, not really
			log.Fatal("FATAL : Error occured. Can't connect to database")
		}
	}
	return mgoSession.Clone()
}

/**
 * get Tweets from DB
 */
func getTweets(ts long) {
	//check that connection openned
	openDbConnection()

	c := session.DB(databaseName).C(collectionName)
	err = c.Find(bson.M{"timestamp": {"$gt": ts}})
	if err != nil {
		log.Fatal("Error getting tweets" + err)
	}
	tweets, _ := json.Marshal(result)
	return
}

/**
 * Add tweet into DB
 */
func addTweet(t string) {
	c := mgoSession.DB(databaseName).C(collectionName)
	err = c.Insert(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Beer %s created\n", beer.Name)
}
*/