package main

import (
	"github.com/ChimeraCoder/anaconda"
	"mhacks2013f/alchemy"
	"mhacks2013f/bass"
	"mhacks2013f/util"
)

func main() {
	alchemy_credentials, _ := util.GenFileHash("credentials/alchemy_credentials")
	twitter_credentials, _ := util.GenFileHash("credentials/twitter_credentials")

	anaconda.SetConsumerKey(twitter_credentials["CONSUMER_KEY"])
	anaconda.SetConsumerSecret(twitter_credentials["CONSUMER_SECRET"])
	api := anaconda.NewTwitterApi(twitter_credentials["ACCESS_TOKEN"], twitter_credentials["ACCESS_TOKEN_SECRET"])
	
	c := bass.NewRDB()
	countries := bass.GetCountries(c)
	for _, country := range countries {
		var sentiment_val float64
		var pos_words []string
		var neg_words []string
		searchResult, _ := api.GetSearch(country, nil)
		for _ , tweet := range searchResult {
			x,y,z := alchemy.GetSentimentalWords(alchemy_credentials, tweet.Text)
			sentiment_val += x
			pos_words = append(pos_words, y...)
			neg_words = append(neg_words, z...)
		}
		bass.PushRDB(c, country, sentiment_val, pos_words, neg_words)
	}

	bass.CloseRDB(c)
}
