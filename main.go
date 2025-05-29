package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	repo "quote_book/repository"
	"strconv"
	"time"
)

func main() {
	StartHttpServer()
}

func StartHttpServer() {
	repo.Initialize_repo()
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 32,
	}
	http.HandleFunc("GET /quotes", GET_quotes_controller)
	http.HandleFunc("POST /quotes", POST_handler)
	http.HandleFunc("GET /quotes/random", GET_RANDOM)
	http.HandleFunc("DELETE /quotes/", DELETE_by_id)
	log.Fatal(server.ListenAndServe())
}

func GET_quotes_controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Has("author") {
			GET_BY_AUTHOR_handler(w, r)
		} else {
			GET_ALL_handler(w, r)
		}
	default:
		http.Error(w, "invalid method", 400)
	}
}

func GET_ALL_handler(w http.ResponseWriter, r *http.Request) {
	result := repo.GetAll()
	byte_result, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "invalid string", 400)
	}
	w.Write(byte_result)
}

func GET_BY_AUTHOR_handler(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	result := repo.GetAllByAuthor(author)
	byte_result, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
	}
	w.Write(byte_result)

}

func POST_handler(w http.ResponseWriter, r *http.Request) {
	var entry repo.Quote
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		log.Print(err)
		http.Error(w, "invalid input", 400)
	}
	repo.AddEntry(entry)
	w.Write([]byte("entry added"))
}

func GET_RANDOM(w http.ResponseWriter, r *http.Request) {
	result := repo.GetRandom()
	byte_result, err := json.Marshal(result)
	if err != nil {
		log.Print(err)
	}
	w.Write(byte_result)
}

func DELETE_by_id(w http.ResponseWriter, r *http.Request) {
	path := path.Base(r.URL.Path)
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "invalid id", 400)
	}
	err = repo.RemoveAtIndex(id)
	if err != nil {
		http.Error(w, "invalid id", 400)
	}
	w.Write([]byte("entry deleted"))

}
