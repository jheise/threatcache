package main

import (
    "fmt"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<html><head></head><body>/ip<br/>/domain<br/>/email<br/>/file<br/>/antivirus<br/></body></html>", r.URL.Path[1:])
}

func IPHandler(w http.ResponseWriter, r *http.Request){
    ipaddr := r.URL.Query().Get("ip")
    fmt.Printf("Fetching %s\n", ipaddr)
    fmt.Fprintf(w, "stuff", r.URL.Path[1:])
}

func main(){
    fmt.Println("Listening...")
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/ip", IPHandler)
    http.ListenAndServe(":8888", nil)

}
