package session

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type redisStore struct {
	client *redis.Client
}

func NewRedisStore() Store {
	//instantiate redis object
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to ping Redis: %v", err)
	}
	log.Println("Connected to Redis Successfully")

	return &redisStore{
		client: client,
	}
}

func (r redisStore) Get(id string) (SessionToken, error) {
	var session SessionToken

	body, err := r.client.Get(id).Bytes()
	if err != nil {
		fmt.Printf("failed to save session to redis, %v", err)
		return session, err
	}

	if err := json.Unmarshal(body, &session); err != nil {
		fmt.Printf("failed to save session to redis, %v", err)
		return session, err
	}

	return session, nil
}

func (r redisStore) Set(id string, session SessionToken) error {
	body, err := json.Marshal(session)
	if err != nil {
		log.Printf("failed to save session to redis, %v", err)
	}
	//set expire time the same with cookie expiration
	err = r.client.Set(id, body, time.Hour*10).Err()
	if err != nil {
		log.Printf("failed to save session to redis, %v", err)
	}
	return nil
}

func (r redisStore) Del(id string) error {
	err := r.client.Del(id).Err()
	if err != nil {
		log.Printf("failed to delete session to redis, %v", err)
	}
	return nil
}
