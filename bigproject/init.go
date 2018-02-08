package bigproject

import (
	"log"

	"github.com/tokopedia/gosample/bigproject/src/database"
)

func InitBP() {
	err := database.InitDB()
	if err != nil {
		log.Println(err)
	}

	database.InitRedis()
}
