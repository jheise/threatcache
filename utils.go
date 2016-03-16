package main

import (
    // internal libraries
    "fmt"
    "net/http"

    // external libraries
    "github.com/garyburd/redigo/redis"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<html><head></head><body>/ip<br/>/domain<br/>/email<br/>/file<br/>/antivirus<br/></body></html>")
}

func GetKey(key string) (string, error){
    rclient, err := redis.Dial("tcp", RedisConnect)
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
    rclient, err := redis.Dial("tcp", RedisConnect)
    if err != nil {
        return err
    }
    defer rclient.Close()

    _, err = rclient.Do("SET", key, value)
    if err != nil {
        return err
    }

    //Set data to expire
    _, err = rclient.Do("EXPIRE", key, Timeout)
    return err
}
