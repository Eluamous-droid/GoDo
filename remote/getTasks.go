package remote

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetAllItems() []TodoItem {

	resp, err := http.Get("http://localhost:8080/getAllItems")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var items []TodoItem
	if err != json.Unmarshal(body, &items) {
		panic(err)
	}

	return items
}

func GetAllItemsInGroup(group string) []TodoItem {

	url := "http://localhost:8080/getAllItemsInGroup"
	var jsonInput = []byte(`{"group":"` + group + `"}`)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var items []TodoItem
	if err != json.Unmarshal(body, &items) {
		panic(err)
	}

	return items
}
