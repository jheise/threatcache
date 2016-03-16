package main

import (
    // internal libraries
    "fmt"
    "net/http"

    // external libraries
    "github.com/gorilla/mux"
    "github.com/jheise/gothreat"
)

func AVHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    av := vars["av"]
    Info.Printf("Fetching %s\n", av)
    key := fmt.Sprintf("av - %s", av)
    data, err := GetKey(key)
    if err != nil {
        Info.Println(err)
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
