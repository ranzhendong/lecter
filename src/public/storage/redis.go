package storage

import (
	"fmt"
	"github.com/go-redis/redis"
)

func redisconnect(ipPort, passWord string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     ipPort,
		Password: passWord, // no password set
		DB:       0,        // use default DB
	})

	_, err := client.Ping().Result()
	//fmt.Printf("%T", client)
	//fmt.Println(pong, err)
	if err != nil {
		return client
	}
	//fmt.Println(pong, err)
	return client
	// Output: PONG <nil>
}

func RedisReset(ipPort, passWord string) {
	client := redisconnect(ipPort, passWord)
	_ = client.FlushAll()
	client.Close()
}

func RedisSet(ipPort, passWord, name string, bodyContent interface{}) {
	//fmt.Println("RedisSet:   ", bodyContent)
	client := redisconnect(ipPort, passWord)
	err := client.Set(name, bodyContent, 0).Err()
	client.Close()
	if err != nil {
		panic(err)
	}
}

func RedisGet(ipPort, passWord, name string) (val string) {
	client := redisconnect(ipPort, passWord)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		client.Close()
	}()
	val, err := client.Get(name).Result()
	client.Close()
	if err != nil {
		panic(err)
	}
	return val
}
