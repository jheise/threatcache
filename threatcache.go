package main

import (
    // internal libraries
    "flag"
    "log"
    "net/http"
    "os"

    // external libraries
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

var (
    Timeout string      // Time in seconds before cache eviction
    RedisConnect string // host:port to use in redis connection
    ListenInfo string   // IP address to listen on
    Info    *log.Logger // Logger
)

func init(){
    var redishost string
    var redisport string
    var hostip string
    var hostport string

    flag.StringVar(&redishost, "redishost", "localhost", "Redis host to connect to")
    flag.StringVar(&redisport, "redisport", "6379", "Redis port to connect on")
    flag.StringVar(&Timeout, "timeout", "14400", "TTL for caching")
    flag.StringVar(&hostip, "ipaddr", "0.0.0.0", "Address to bind on")
    flag.StringVar(&hostport, "port", "8888", "Port to bind on")
    flag.Parse()

    RedisConnect = redishost + ":" + redisport
    ListenInfo = hostip + ":" + hostport
}

func main(){
    Info = log.New(os.Stdout, "INFO: ",
            log.Ldate|log.Ltime|log.Lshortfile)

    r := mux.NewRouter()
    r.HandleFunc("/", IndexHandler)
    r.HandleFunc("/ip/{ipaddr}/", IPHandler)
    r.HandleFunc("/domain/{domain}/", DomainHandler)
    r.HandleFunc("/email/{email}/", EmailHandler)
    r.HandleFunc("/file/{file}/", FileHandler)
    r.HandleFunc("/av/{av}/", AVHandler)
    Info.Println("Listening...")
    http.Handle("/", r)
    loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, r)
    http.ListenAndServe(ListenInfo, loggedRouter)
}
