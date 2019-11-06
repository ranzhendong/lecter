package lepai_storage

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
	err := client.FlushAll()
	fmt.Println(err)
}

func RedisSet(ipPort, passWord, name string, bodyContent interface{}) {
	//fmt.Println("RedisSet:   ", bodyContent)
	client := redisconnect(ipPort, passWord)
	err := client.Set(name, bodyContent, 0).Err()
	if err != nil {
		panic(err)
	}
}

func RedisGet(ipPort, passWord, name string) (val string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	client := redisconnect(ipPort, passWord)
	val, err := client.Get(name).Result()
	if err != nil {
		panic(err)
	}
	return val
}
