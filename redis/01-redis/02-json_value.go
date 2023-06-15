package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Author{Name: "Mostafa", Age: 25})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(ctx, "id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(ctx, "id1234").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
