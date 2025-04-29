package pr

import (
	"encoding/json"
	"log"
)

func JSON(v any) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Println("Invalid JSON")
	}

	log.Println(string(b))
}
