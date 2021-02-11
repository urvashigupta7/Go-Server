package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func homePage(w http.ResponseWriter, r *http.Request){
   var randomString [5]string
   todos, err := db.Query("SELECT * FROM test")
    if err != nil {
        panic(err.Error())
    }
    for todos.Next(){
        var id int
        var todo string
        err = todos.Scan(&id,&todo)
        if err != nil {
            panic(err.Error())
        }
        randomString[id]=todo
    }
   json.NewEncoder(w).Encode(randomString)
}

func handleRequests() {
    router := mux.NewRouter()
    router.HandleFunc("/todo",homePage)
    log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
    db,err= sql.Open("mysql","urgupta:urgupta@tcp(mysqldb:3306)/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    _,err = db.Exec("CREATE TABLE IF NOT EXISTS test ( id integer, todo varchar(52) )")
    if err != nil {
       panic(err)
    }
    randomString := []string{
        "Learn React",
        "Use Typescript with React",
        "Have Fun!!",
        "Use emoji's",
        "Look at memes",
    }
    for id, todo := range randomString {
        insert,err := db.Prepare("INSERT INTO test(id, todo) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insert.Exec(id,todo)
        log.Println("todo: " + todo)
    }
    
    handleRequests()
}