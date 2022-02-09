package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// User is the object stored in Redis
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    string `json:"id"`
}

var user User
var redisHost = os.Getenv("REDIS_HOST")
var client = redis.NewClient(&redis.Options{
	Addr:     redisHost + ":6379",
	Password: "",
	DB:       0,
})

// handePost handles HTTP POST requests
func handlePost(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = client.Set(user.Id, json, 0).Err()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Storing data: ", string(json))

}

// handleGet handles HTTP GET requests
func handleGet(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	val, err := client.Get(user.Id).Result()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, val)
	log.Println("Retrieving data: ", val)
}

func main() {

	defer client.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", handlePost).Methods("POST")
	r.HandleFunc("/", handleGet).Methods("GET")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(srv.ListenAndServe())

}
