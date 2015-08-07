package main

import (
    "log"
    "fmt"
    "time"
    "net/http"
    "github.com/darkhelmet/twitterstream"
)

var (
        consumerKey = "sc36rGrawG8TlCap6D0moB7BH"
        consumerSecret = "sKvz1idjisqbWxv1hhQNzpFyuzLPba7SCMf1vYmwoy9qTlGmpi"
        accessToken = "2302871682-eAyD8irASZOmqZhrCrMXuWvzPaNq2bzQXKvMyLr"
        accessSecret = "zp3XtajLNXGf72l9YSFdgCDgNaH97O8unlurYOSgG70zD"
)



func decode(conn *twitterstream.Connection) {
    for {
        if tweet, err := conn.Next(); err == nil {
                fmt.Println("%s said: %s", tweet.User.ScreenName, tweet.Text)
            //TODO store in mongo
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
            conn, err := client.Track("#LOSCPSG")//os.Getenv("following")
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

    /*
    consumerKey := os.Getenv("consumerKey")
    consumerSecret := os.Getenv("consumerSecret")
    accessToken := os.Getenv("accessToken")
    accessSecret := os.Getenv("accessSecret")
    */


}

