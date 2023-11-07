package main

import (
	"fmt"

	"github.com/stanley-tarce/webhook-go/redis"
)
func main() {

	conn, err := redis.ConnectRedis()
	if err != nil {
		fmt.Println(err)
		return
	}
	

	fmt.Printf("conn: %v\n", conn)
	
}

