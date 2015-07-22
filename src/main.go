package main

import (
    "log"
    "time"
    "net/http"
    "github.com/darkhelmet/twitterstream"
    "os"
)

func decode(conn *twitterstream.Connection) {
    for {
        if tweet, err := conn.Next(); err == nil {
            log.Println("%s said: %s", tweet.User.ScreenName, tweet.Text)
            //TODO store in mongo

	} else {
            log.Printf("Failed decoding tweet: %s", err)
            return
        }
    }
}

func main() {
	log.Println("Starting app.\n")
    //start webserver
    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))

    //start stream listener
	client := twitterstream.NewClient(os.Getenv("consumerKey"), os.Getenv("consumerSecret"),
            os.Getenv("accessToken"), os.Getenv("accessSecret"))
    for {
        conn, err := client.Follow("CMBretagne")
        if err != nil {
            log.Println("Tracking failed, sleeping for 1 minute")
            time.Sleep(1 * time.Minute)
            continue
        }
        decode(conn)
    }
}

