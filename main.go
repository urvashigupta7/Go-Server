package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
type test_struct struct {
    Todo string
}
func homePage(w http.ResponseWriter, r *http.Request){

switch r.Method {
case "GET":
    numberOfTodos,err := db.Query("SELECT COUNT(*) AS nos FROM test")
    if err != nil {
        panic(err.Error())
    }
    var nos int
    numberOfTodos.Next()
    err = numberOfTodos.Scan(&nos)
    if err != nil {
        panic(err.Error())
    }
    randomString := make([]string, nos+1)
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
case "POST":
   decoder := json.NewDecoder(r.Body)
   var t test_struct
   err := decoder.Decode(&t)
   if err != nil {
    panic(err.Error())
   }
   insert,err := db.Prepare("INSERT INTO test(todo) VALUES(?)")
    if err != nil {
        panic(err.Error())
    }
    insert.Exec(t.Todo)
    fmt.Println(t.Todo)
default:
    fmt.Println("only get and post request supported")
}
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
    _,err = db.Exec("CREATE TABLE IF NOT EXISTS test ( id integer NOT NULL AUTO_INCREMENT, todo varchar(52),  PRIMARY KEY (id) )")
    if err != nil {
       panic(err)
    }
    handleRequests()
}

// "Learn React",
// "Use Typescript with React",
// "Have Fun!!",
// "Use emoji's",
// "Look at memes",