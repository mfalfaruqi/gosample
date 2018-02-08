package database

import (
	"log"
)

func IncVisitorCount() (int, error) {
	val, err := client.Incr("visitor_count").Result()

	if err != nil {
		log.Println(err)
	}

	return int(val), nil
}
