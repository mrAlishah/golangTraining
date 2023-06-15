package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// we can call set with a `Key` and a `Value`.
	err = client.Set("name", "Mostafa", 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
