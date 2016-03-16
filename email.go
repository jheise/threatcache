package main

import (
    // internal libraries
    "fmt"
    "net/http"

    // external libraries
    "github.com/gorilla/mux"
    "github.com/jheise/gothreat"
)

func EmailHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    email := vars["email"]
    Info.Printf("Fetching %s\n", email)
    key := fmt.Sprintf("email - %s", email)
    data, err := GetKey(key)
    if err != nil {
        Info.Println(err)
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
