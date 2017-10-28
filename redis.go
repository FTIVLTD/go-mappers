package mappers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// Redis - redis DB mapper
type Redis struct {
	DBConfig
	client *redis.Client
}

/*
Get - getting data by key
*/
func (r *Redis) Get(key string) (interface{}, error) {
	err := r.checkConnection()
	if err != nil {
		return nil, err
	}
	// TODO: check for redis.Nil: if err == redis.Nil {
	return r.client.Get(key).Result()
}

func (r *Redis) Del(string) error {
	return nil
}

/*
Set - setting key with value and expiration time
*/
func (r *Redis) Set(key string, value interface{}, exp time.Duration) error {
	err := r.checkConnection()
	if err != nil {
		return err
	}

	str, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(key, str, exp).Err()
}

func (r *Redis) connect() error {
	r.client = redis.NewClient(&redis.Options{
		Addr:     r.getAddr(),
		Password: r.Password, // no password set
		DB:       0,          // use default DB
	})
	_, err := r.client.Ping().Result()
	return err
}

func (r *Redis) checkConnection() error {
	if r.client == nil {
		return r.connect()
	}
	return nil
}

func (r *Redis) getAddr() string {
	return r.Host + ":" + strconv.Itoa(r.Port)
}
