package cats_api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type catjson struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func ApiCall(url string) []byte {

	req, errget := http.Get(url)

	if errget != nil {
		log.Fatalf("An Error occured! err: %s", errget)
	}

	defer req.Body.Close()

	response, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatalf("An Error occured! err: %s", err)
	}

	return response

}

func JsonUnmarshall(response []byte) catjson {

	var response_json catjson

	json.Unmarshal(response, &response_json)

	return response_json

}
