package main

import (
    "log"
    "fmt"
    "time"
    "net/http"
    "github.com/darkhelmet/twitterstream"
)

var (
        consumerKey = os.Getenv("consumerKey")
        consumerSecret = os.Getenv("consumerSecret")
        accessToken = os.Getenv("accessToken")
        accessSecret = os.Getenv("accessSecret")
)



func decode(conn *twitterstream.Connection) {
    for {
        if tweet, err := conn.Next(); err == nil {
            addTweet(tweet)
	} else {
            fmt.Println("Failed decoding tweet: %s", err)
            log.Printf("Failed decoding tweet: %s", err)
            return
        }
    }
}

func startStreaming() {
    //start stream listener
    fmt.Println("start twitter stream", consumerKey, consumerSecret, accessToken, accessSecret)
    client := twitterstream.NewClient(consumerKey, consumerSecret, accessToken, accessSecret)
    for {
        conn, err := client.Track(os.Getenv("hashtag"))
        if err != nil {
            fmt.Println("Tracking failed, sleeping for 1 minute %s", err)
            log.Println("Tracking failed, sleeping for 1 minute")
            time.Sleep(1 * time.Minute)
            continue
        }
        decode(conn)
    }
}

func main() {
	log.Println("Starting app.\n")
    //start webserver
    router := NewRouter()

    go startStreaming()

    log.Fatal(http.ListenAndServe(":8080", router))
}

