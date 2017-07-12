package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    fmt.Println("Listening on port 8080")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {   	
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}

// package main

// import (
//     "fmt"
//     "log"
//     "net/http"

//     "github.com/gorilla/mux"
// )

// func main() {

//     router := mux.NewRouter().StrictSlash(true)
//     router.HandleFunc("/", Index)
//     router.HandleFunc("/todos", TodoIndex)
//     router.HandleFunc("/todos/{todoId}", TodoShow)

//     fmt.Println("Listening on port 8080")

//     log.Fatal(http.ListenAndServe(":8080", router))
// }

// func Index(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "Welcome!")
// }

// func TodoIndex(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "Todo Index!")
// }

// func TodoShow(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     todoId := vars["todoId"]
//     fmt.Fprintln(w, "Todo show:", todoId)
// }