package main

import (
	"fmt"
	"log"
	"net/http"
	"paddycache/paddycache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	paddycache.NewGroup("scores", 2<<10, paddycache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "localhost:9999"
	peers := paddycache.NewHTTPPool(addr)
	log.Println("paddycache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
