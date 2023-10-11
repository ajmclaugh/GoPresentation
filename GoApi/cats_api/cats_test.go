package cats_api

import (
	"testing"
)

func TestApicall(t *testing.T) {
	url := "https://catfact.ninja/fact"

	response := ApiCall(url)

	response_json := JsonUnmarshall(response)
	if response_json.Length == 0 {
		t.Fatalf("Lenght is %v", response_json.Length)
	}

}
