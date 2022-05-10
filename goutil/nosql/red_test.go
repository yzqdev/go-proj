package nosql

import (
	"github.com/fatih/color"
	"github.com/go-redis/redis"
	"testing"
)

func TestRed(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	color.Red("hhhhhhh")
	color.Cyan(pong,err)
}