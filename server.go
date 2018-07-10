package main

import (
    "encoding/json"
    "log"
    "net/http"
    "io/ioutil"
)

/*type test_struct struct {
    Test string
}*/

func webhook(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
    /*var t test_struct
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
    }
    log.Println(t.Test)*/
}

func main() {
    http.HandleFunc("/webhook", webhook)
    log.Fatal(http.ListenAndServe(":8080", nil))
}