package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	repo "quote_book/repository"
	"testing"
)

var addr string = "http://localhost:8080/quotes"

func TestConnect(t *testing.T) {
	_, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func TestAddEntry(t *testing.T) {
	entry := repo.Quote{Author: "Confucius", Quote: "Life is simple, but we insist on making it complicated."}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(entry)
	_, err := http.Post(addr, "application/json", &buf)
	if err != nil {
		log.Fatal(err)
	}

}

func TestFetchEntries(t *testing.T) {
	response, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body_bytes, _ := io.ReadAll(response.Body)
	var parsed_result []repo.Quote
	json.Unmarshal(body_bytes, &parsed_result)

	if parsed_result[0].Author != "Confucius" || parsed_result[0].Quote != "Life is simple, but we insist on making it complicated." {
		t.Errorf("want, Confucius" + "got " + string(body_bytes))
	}
}
