package main

import (
    // internal libraries
    "fmt"
    //"log"
    "net/http"
    "os"

    // external libraries
    "github.com/garyburd/redigo/redis"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "github.com/jheise/gothreat"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<html><head></head><body>/ip<br/>/domain<br/>/email<br/>/file<br/>/antivirus<br/></body></html>")
}

func IPHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    ipaddr := vars["ipaddr"]
    fmt.Printf("Fetching %s\n", ipaddr)
    key := fmt.Sprintf("ip - %s", ipaddr)
    data, err := GetKey(key)
    if err != nil {
        fmt.Println(err)
    }

    // if data is length 0 there is no record, fetch from source
    if len(data) > 0 {
        fmt.Fprintf(w, "%s", data)
    } else {
        ipdata, err := gothreat.IPReportRaw(ipaddr)
        if err != nil {
            fmt.Fprintf(w, "Error processing %s", ipaddr)
        } else {
            err = SetKey(key, fmt.Sprintf("%s", ipdata))
            if err != nil {
                fmt.Fprintf(w, "Error Storing %s", ipaddr)
            } else {
                fmt.Fprintf(w, "%s", ipdata)
            }
        }
    }
}

func DomainHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    domain := vars["domain"]
    fmt.Printf("Fetching %s\n", domain)
    key := fmt.Sprintf("domain - %s", domain)
    data, err := GetKey(key)
    if err != nil {
        fmt.Println(err)
    }

    // if data is length 0 there is no record, fetch from source
    if len(data) > 0 {
        fmt.Fprintf(w, "%s", data)
    } else {
        domaindata, err := gothreat.DomainReportRaw(domain)
        if err != nil {
            fmt.Fprintf(w, "Error processing %s", domain)
        } else {
            err = SetKey(key, fmt.Sprintf("%s", domaindata))
            if err != nil {
                fmt.Fprintf(w, "Error Storing %s", domain)
            } else {
                fmt.Fprintf(w, "%s", domaindata)
            }
        }
    }
}

func EmailHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    email := vars["email"]
    fmt.Printf("Fetching %s\n", email)
    key := fmt.Sprintf("email - %s", email)
    data, err := GetKey(key)
    if err != nil {
        fmt.Println(err)
    }

    // if data is length 0 there is no record, fetch from source
    if len(data) > 0 {
        fmt.Fprintf(w, "%s", data)
    } else {
        emaildata, err := gothreat.EmailReportRaw(email)
        if err != nil {
            fmt.Fprintf(w, "Error processing %s", email)
        } else {
            err = SetKey(key, fmt.Sprintf("%s", emaildata))
            if err != nil {
                fmt.Fprintf(w, "Error Storing %s", email)
            } else {
                fmt.Fprintf(w, "%s", emaildata)
            }
        }
    }
}

func FileHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    file := vars["file"]
    fmt.Printf("Fetching %s\n", file)
    key := fmt.Sprintf("file - %s", file)
    data, err := GetKey(key)
    if err != nil {
        fmt.Println(err)
    }

    // if data is length 0 there is no record, fetch from source
    if len(data) > 0 {
        fmt.Fprintf(w, "%s", data)
    } else {
        filedata, err := gothreat.FileReportRaw(file)
        if err != nil {
            fmt.Fprintf(w, "Error processing %s", file)
        } else {
            err = SetKey(key, fmt.Sprintf("%s", filedata))
            if err != nil {
                fmt.Fprintf(w, "Error Storing %s", file)
            } else {
                fmt.Fprintf(w, "%s", filedata)
            }
        }
    }
}

func AVHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    av := vars["av"]
    fmt.Printf("Fetching %s\n", av)
    key := fmt.Sprintf("av - %s", av)
    data, err := GetKey(key)
    if err != nil {
        fmt.Println(err)
    }

    // if data is length 0 there is no record, fetch from source
    if len(data) > 0 {
        fmt.Fprintf(w, "%s", data)
    } else {
        avdata, err := gothreat.FileReportRaw(av)
        if err != nil {
            fmt.Fprintf(w, "Error processing %s", av)
        } else {
            err = SetKey(key, fmt.Sprintf("%s", avdata))
            if err != nil {
                fmt.Fprintf(w, "Error Storing %s", av)
            } else {
                fmt.Fprintf(w, "%s", avdata)
            }
        }
    }
}

func GetKey(key string) (string, error){
    rclient, err := redis.Dial("tcp", ":6379")
    if err != nil {
        return "",err
    }
    defer rclient.Close()

    data, err := redis.String(rclient.Do("GET", key))
    if err != nil {
        return "", err
    }
    return data, err
}

func SetKey(key string, value string) error {
    rclient, err := redis.Dial("tcp", ":6379")
    if err != nil {
        return err
    }
    defer rclient.Close()

    _, err = rclient.Do("SET", key, value)
    if err != nil {
        return err
    }

    //Set data to expire
    _, err = rclient.Do("EXPIRE", key, "10")
    return err
}


func main(){

    r := mux.NewRouter()
    r.HandleFunc("/", IndexHandler)
    r.HandleFunc("/ip/{ipaddr}/", IPHandler)
    r.HandleFunc("/domain/{domain}/", DomainHandler)
    r.HandleFunc("/email/{email}/", EmailHandler)
    r.HandleFunc("/file/{file}/", FileHandler)
    r.HandleFunc("/av/{av}/", AVHandler)
    fmt.Println("Listening...")
    http.Handle("/", r)
    loggedRouter := handlers.LoggingHandler(os.Stdout, r)
    http.ListenAndServe(":8888", loggedRouter)
}
