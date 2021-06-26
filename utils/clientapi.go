package clientapi

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetJson(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var parsedBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&parsedBody)
	defer resp.Body.Close()

	return parsedBody
}
