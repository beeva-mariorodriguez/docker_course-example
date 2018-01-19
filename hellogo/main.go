package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
    "github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

const version = "0.4"
var rc *redis.Client

func main() {
	log.Println("starting hellogo ...")
    rc = redisClient()
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/health", health).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent())
	w.WriteHeader(http.StatusOK)
	h, _ := os.Hostname()
	fmt.Fprintf(w, "%s > hello GO %s\n", h, version)
}

func health(w http.ResponseWriter, r *http.Request) {
    _, err := rc.Ping().Result()
    if err != nil{
        log.Println(r.UserAgent())
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "KO: redis: ")
        fmt.Fprintln(w, err)
    }else{
        log.Println(r.UserAgent())
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "OK")
    }
}

func redisClient() *redis.Client {
    redisdb,_:=strconv.Atoi(os.Getenv("REDISDB"))
    redishost:=os.Getenv("REDISHOST")
    redisport:=os.Getenv("REDISPORT")

	return redis.NewClient(&redis.Options{
		Addr: redishost + ":" + redisport,
		DB: redisdb,
		Password: "",
	})
}
