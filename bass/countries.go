package bass

import (
	"encoding/json"
	"github.com/simonz05/godis/redis"
	"log"
	"time"
)

type CountrySentiment struct {
	Timestamp time.Time `json:"timestamp"`
	Values float64 `json:"values"`
	Pos []string `json:"pos"`
	Neg []string `json:"neg"`
}

func GetCountries(c *redis.Client) []string {
	reply, _ := c.Smembers("countries")
	str := reply.StringArray()
	return str
}

func PushRDB(c *redis.Client, country string, sentiment_val float64, pos_words, neg_words []string) {
	cs := CountrySentiment{Timestamp: time.Now(), Values: sentiment_val, Pos: pos_words, Neg: neg_words}
	b, err := json.Marshal(cs)
	if err != nil {
		log.Fatal(err)
	}
	c.Sadd(country, b)
}
