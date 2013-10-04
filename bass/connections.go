package bass

import (
	"github.com/simonz05/godis/redis"
	"log"
	"os"
	"strconv"
	"mhacks2013f/util"
)

func getEnvironment() string {
	env := os.Getenv("RAILS_ENV")
	switch env {
	case "development", "production":
		return env
	}
	return "development"
}

func NewRDB() *redis.Client {
	rdb_info, err := util.GenFileHash("environments/" + getEnvironment() + "_redis")
	rdb_num, _ := strconv.Atoi(rdb_info["NUM"])
	if err != nil {
		log.Fatal(err)
	}
	return redis.New("tcp:"+rdb_info["HOST"]+":"+rdb_info["PORT"], rdb_num, rdb_info["PASS"])
}

func PullRDB(country_name string, c *redis.Client) []string {
	reply, _ := c.Smembers(country_name)
	str := reply.StringArray()
	return str
}

func CloseRDB(c *redis.Client) {
	c.Quit()
}
