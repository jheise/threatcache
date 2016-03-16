package main

import (
    // internal libraries
    "fmt"
    "net/http"

    // external libraries
    "github.com/gorilla/mux"
    "github.com/jheise/gothreat"
)

func DomainHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    domain := vars["domain"]
    Info.Printf("Fetching %s\n", domain)
    key := fmt.Sprintf("domain - %s", domain)
    data, err := GetKey(key)
    if err != nil {
        Info.Println(err)
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
