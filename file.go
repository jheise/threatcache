package main

import (
    // internal libraries
    "fmt"
    "net/http"

    // external libraries
    "github.com/gorilla/mux"
    "github.com/jheise/gothreat"
)

func FileHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    file := vars["file"]
    Info.Printf("Fetching %s\n", file)
    key := fmt.Sprintf("file - %s", file)
    data, err := GetKey(key)
    if err != nil {
        Info.Println(err)
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
