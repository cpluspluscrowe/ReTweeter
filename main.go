package main

import (
	"fmt"
	"github.com/amit-lulla/twitterapi"
	"github.com/deckarep/golang-set"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
	"os"
	"time"
)

var tweetIds = mapset.NewSet()

func doEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

func main() {
	doEvery(200000*time.Millisecond, Retweet)
}

func Retweet() {
	tweetIds = SearchAndFavorite("Kotlin", tweetIds)
	tweetIds = SearchAndFavorite("Golang", tweetIds)
	tweetIds = SearchAndFavorite("Scala", tweetIds)
	tweetIds = SearchAndFavorite("Spark", tweetIds)
	tweetIds = SearchAndFavorite("Python", tweetIds)
	fmt.Println()
}

func SearchAndFavorite(look4 string, tweetIds mapset.Set) mapset.Set {
	twitterClient := getTwitterClient()
	search, _, err := twitterClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: look4,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Searching")
	for _, tweet := range search.Statuses {
		tweet_id := tweet.ID
		//text := tweet.Text
		//twitterClient.Statuses.Retweet(tweet_id, &twitter.StatusRetweetParams{})
		if !tweetIds.Contains(tweet_id) {
			fmt.Println(tweet_id)
			//Favorite(tweet_id)
		}
		tweetIds.Add(tweet_id)
	}
	return tweetIds
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

func Favorite(tweetId int64) {
	consumerKey := os.Getenv("consumerKey")
	consumerSecret := os.Getenv("consumerSecret")
	accessToken := os.Getenv("accessToken")
	accessSecret := os.Getenv("accessSecret")

	twitterapi.SetConsumerKey(consumerKey)
	twitterapi.SetConsumerSecret(consumerSecret)
	api := twitterapi.NewTwitterApi(accessToken, accessSecret)
	_, err := api.Favorite(tweetId)
	if err != nil {
		fmt.Println(err)
		return
	}
}
