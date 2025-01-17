package utils

import (
	"encoding/json"
	"log"
)

func PrintStruct[T any](obj T) {
	jsonData, err := json.MarshalIndent(obj, "", "   ")
	if err != nil {
		log.Printf("PrintStruct Method error : %v", err)
		return
	}

	log.Printf("Debug Ojb : %v", string(jsonData))
}
