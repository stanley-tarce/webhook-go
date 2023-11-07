package redis

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)


func ConnectRedis() (redis.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	rEndpoint := os.Getenv("REDIS_ENDPOINT")
	rPort := os.Getenv("REDIS_PORT")
	rPassword := os.Getenv("REDIS_PASSWORD")
	rTls, strConvErr := strconv.ParseBool(os.Getenv("REDIS_TLS_ENABLED"))
	endpoint := strings.Join([]string{rEndpoint, ":",rPort}, "")

	if strConvErr != nil {
		return nil, strConvErr
	}

	if len(endpoint) == 0 {
		return nil, errors.New("Failed to get environment variable")
	}
	c, err := redis.Dial("tcp", endpoint, redis.DialUseTLS(rTls))
	if err != nil {
		return nil, err
	}
	_, err = c.Do("AUTH", rPassword)
	if err != nil {
		return nil, err
	}

	return c, nil
}