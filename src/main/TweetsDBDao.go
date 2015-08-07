package main
import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"fmt"
	    "github.com/darkhelmet/twitterstream"

)

var (
	mgoSession   *mgo.Session
	databaseName = "tweets"
	collectionName = "tweets"
	databaseUri  = os.Getenv("MONGODB_ADDON_URI")
)

func openDbConnection() *mgo.Session {
	if mgoSession == nil {
		var err error
		fmt.Println("database uri : ", databaseUri)
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
func getTweets(ts int64) string {
	//check that connection openned
	openDbConnection()
	var results []Tweet
	var err error
	c := mgoSession.DB(databaseName).C(collectionName)
	search := bson.M{"timestamp": bson.M{"$gt": ts}}
	err = c.Find(search).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	tweets, _ := json.Marshal(results)
	return string(tweets)
}

/**
 * Add tweet into DB
 */
func addTweet(t *twitterstream.Tweet) {
	fmt.Println(t)
	openDbConnection()
	c := mgoSession.DB(databaseName).C(collectionName)
	err := c.Insert(t)
	if err != nil {
		log.Fatal(err)
	}
}
