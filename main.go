package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)


func homePage(w http.ResponseWriter, r *http.Request){
   randomString := []string{
       "Learn React",
       "Use Typescript with React",
       "Have Fun!!",
       "Use emoji's",
       "Look at memes",
   }
   json.NewEncoder(w).Encode(randomString)
}

func handleRequests() {
    router := mux.NewRouter()
    router.HandleFunc("/todo",homePage)
    log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
    handleRequests()
}