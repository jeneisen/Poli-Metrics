package alchemy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"log"
)

type Sentiment struct {
	Country string `json:"country"`
	Timestamp int64 `json:"timestamp"`
	Positive float64 `json:"positive"`
	Negative float64 `json:"negative"`
	Sentiments []string `json:"sentiments"`
}

func GetSentimentalWords(cred map[string]string, text string) (float64, []string, []string) {
	var sentimental_value float64
	var sentimental_pos_words []string
	var sentimental_neg_words []string
	
	values := make(url.Values)
	values.Set("apikey",cred["ALCHEMIST"])
	values.Set("text",text)
	values.Set("sentiment","1")
	values.Set("outputMode","json")
	
	resp, err := http.PostForm("http://access.alchemyapi.com/calls/text/TextGetRankedKeywords", values)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	type Sentiments struct {
		Type string
		Score string
	}
	type Keywords struct {
		Text string
		Relevance string
		Sentiment Sentiments
	}
	type Analysis struct {
		Status string
		Usage string
		Url string
		Language string
		Keyword []Keywords
	}

	var a Analysis
	err = json.Unmarshal([]byte(string(bs)), &a)
	if err != nil {
		log.Fatal(err)
	}
	j := make(map[string]interface{})
	err = json.Unmarshal([]byte(string(bs)), &j)
	if err != nil {
		log.Fatal(err)
	}
	keywords, _ := json.Marshal(j["keywords"])
	var data []map[string]interface{}
	err = json.Unmarshal([]byte(string(keywords)), &data)
	if err != nil {
		log.Fatal(err)
	}
	for _,y := range data {
		var value float64
		var relevance float64
		var text string
		var text_neutraility int
		for i,j := range y {
			switch i {
			case "sentiment":
				sentiments, _ := json.Marshal(j)
				var sents map[string]string
				err = json.Unmarshal([]byte(string(sentiments)), &sents)
				if err != nil {
					log.Fatal(err)
				}
				switch sents["type"] {
				case "positive":
					text_neutraility = 1
				case "negative":
					text_neutraility = -1
				}
				if sents["type"] != "neutral" {
					value, _ = strconv.ParseFloat(sents["score"],64)
				}
			case "relevance":
				relevance, _ = strconv.ParseFloat(j.(string),64)
			case "text":
				text = j.(string)
			}
		}
		sentimental_value += relevance * value
		switch text_neutraility {
		case 1:
			sentimental_pos_words = append(sentimental_pos_words, text)
		case -1:
			sentimental_neg_words = append(sentimental_neg_words, text)
		}
	}

	return 100*sentimental_value, sentimental_pos_words, sentimental_neg_words
}
