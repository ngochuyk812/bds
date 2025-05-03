package config

import "os"

var (
	Brokers      = os.Getenv("BROKERS_EVENTBUS")
	Topic        = os.Getenv("TOPIC_EVENTBUS")
	DbConnect    = os.Getenv("DB_CONNECTION")
	DbName       = os.Getenv("DB_NAME")
	Passwordort  = os.Getenv("SERVER_PORT")
	Port         = os.Getenv("SERVER_PORT")
	RedisConnect = os.Getenv("REDIS_CONNECTION")
	RedisPass    = os.Getenv("REDIS_PASS")
	SecretKey    = os.Getenv("SECRET_KEY")
)
