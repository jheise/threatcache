package main

import (
    // internal libraries
    "fmt"
    "net/http"

    // external libraries
    "github.com/gorilla/mux"
    "github.com/jheise/gothreat"
)

func IPHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    ipaddr := vars["ipaddr"]
    Info.Printf("Fetching %s\n", ipaddr)
    key := fmt.Sprintf("ip - %s", ipaddr)
    data, err := GetKey(key)
    if err != nil {
        Info.Println(err)
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
