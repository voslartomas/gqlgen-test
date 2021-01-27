package cache

import (
	"bytes"
	"encoding/gob"
	"log"
	"reflect"
	"time"

	"github.com/go-redis/redis"
)

type DecoratorOptions struct {
	Key          string
	TTL          time.Duration
	Interface    interface{}
	IgnoreParams map[int]bool
}

type CachedFunction func(interface{}) (reflect.Value, []byte, error)

func CacheDecorator(originalFunction interface{}, options DecoratorOptions) CachedFunction {
	return func(in interface{}) (reflect.Value, []byte, error) {
		args := []reflect.Value{reflect.ValueOf(in)}
		key := options.Key
		for i := 0; i < len(args); i++ {
			if options.IgnoreParams == nil || !options.IgnoreParams[i] {
				key += args[i].String()
			}
		}

		cached := redisClient.Get(key)
		_, resErr := cached.Result()
		encodedReceived, _ := cached.Bytes()

		if resErr != redis.Nil {
			var v interface{}
			return reflect.ValueOf(v), encodedReceived, nil
		}

		fnType := reflect.ValueOf(originalFunction)
		if fnType.Kind() != reflect.Func {
			panic("Expected a function")
		}

		res := fnType.Call(args)

		err, ok := res[1].Interface().(error)

		if !ok {
			encodedToBeSent := new(bytes.Buffer)
			gob.NewEncoder(encodedToBeSent).Encode(res[0].Interface())

			redisClient.Set(key, encodedToBeSent.Bytes(), options.TTL)
			return res[0], nil, nil
		}

		return res[0], nil, err
	}
}

var redisClient *redis.Client

func Connect() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Println("Redis connected")
}

func GetClient() *redis.Client {
	return redisClient
}
