package redis

import (
	"boilerplate/config"
	"boilerplate/util/logwrapper"
	"fmt"

	"github.com/go-redis/redis/v7"
)

// Connection - Connection structure
type Connection struct {
	Conn     *redis.Client
	Database int
	Done     chan error
}

// Client - Redis Connection
var Client *Connection

// NewConnection - new connection of amqp
func NewConnection(redisConfig config.RedisConfig) error {
	if redisConfig.URL == "" {
		return fmt.Errorf("CONFIGURATION IS MISSING FOR REDIS")
	}

	redisClient := &Connection{
		Conn:     nil,
		Database: redisConfig.Database,
		Done:     make(chan error),
	}

	redisClient.Conn = redis.NewClient(&redis.Options{
		Addr:     redisConfig.URL,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Database, // use default DB
	})
	// defer redisClient.Conn.Close()
	pong, err := redisClient.Conn.Ping().Result()
	if err != nil {
		return err
	}
	logwrapper.Logger.Infoln("Connected to REDIS_URL : ", redisConfig.URL)
	logwrapper.Logger.Infoln("Ping REDIS : ", pong)
	Client = redisClient
	return nil
}
