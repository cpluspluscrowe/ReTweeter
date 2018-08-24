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

func doEvery(d time.Duration, f func(string, *mapset.Set), look4 string, tweetIds *mapset.Set) {
	for _ = range time.Tick(d) {
		f(look4, tweetIds)
	}
}

func main() {
	doEvery(3120*time.Second, SearchAndFavorite, "#Golang", &tweetIds)
	doEvery(4120*time.Second, SearchAndFavorite, "#Kotlin", &tweetIds)
	doEvery(4510*time.Second, SearchAndFavorite, "#Scala", &tweetIds)
	doEvery(5120*time.Second, SearchAndFavorite, "#Java", &tweetIds)
	doEvery(55120*time.Second, SearchAndFavorite, "#Haskell", &tweetIds)
}

func SearchAndFavorite(look4 string, tweetIds *mapset.Set) {
	twitterClient := getTwitterClient()
	search, _, err := twitterClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: look4,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Searching for: ", look4, " number of tweets found: ", len(search.Statuses))
	for _, tweet := range search.Statuses {
		tweet_id := tweet.ID
		if !(*tweetIds).Contains(tweet_id) {
			Favorite(tweet_id)
		} else {
			fmt.Println("tweetIds already contains this tweet!")
		}
		(*tweetIds).Add(tweet_id)
		return
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
	}
}
