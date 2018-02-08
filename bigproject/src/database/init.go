package database

import (
	"github.com/go-redis/redis"
	"github.com/tokopedia/sqlt"
)

var dbUser *sqlt.DB
var client *redis.Client

var err error

// InitDB for init DB connection
func InitDB() error {
	dsn := "host=devel-postgre.tkpd port=5432 user=st140804 password=apaajadeh dbname=tokopedia-user sslmode=disable"
	dbUser, err = sqlt.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return nil
}

// InitRedis for init redis connection
func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
