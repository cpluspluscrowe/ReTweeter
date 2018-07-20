package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
	"os"
	"strings"
	"time"
)

func doEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

func main() {
	doEvery(500000*time.Millisecond, Retweet)
}

func Retweet() {
	Search("#Golang")
}

func Search(look4 string) {
	twitterClient := getTwitterClient()
	search, _, err := twitterClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: look4,
	})
	if err != nil {
		panic(err)
	}
	for _, tweet := range search.Statuses {
		tweet_id := tweet.ID
		text := tweet.Text
		if strings.Contains(text, "#Golang") {
			twitterClient.Statuses.Retweet(tweet_id, &twitter.StatusRetweetParams{})
		}
		break
	}
}

func getClient() *http.Client {
	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")
	accessToken := os.Getenv("accessToken")
	accessSecret := os.Getenv("accessSecret")
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return httpClient
}

func getTwitterClient() *twitter.Client {
	httpClient := getClient()
	client := twitter.NewClient(httpClient)
	return client
}
