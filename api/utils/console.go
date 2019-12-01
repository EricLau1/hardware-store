package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func Debug(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}
