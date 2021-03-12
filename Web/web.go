package main

import (
  "fmt"
  "net/http"

)

func handler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","text/html")
 fmt.Fprint(w,"<h1>Going......!!</h1>")
  fmt.Fprint(w,"<p>Go is fast</p>")

}

func handler_a(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w,"Hey it is also Working........!")
}

func main (){
   http.HandleFunc("/",handler)
   http.HandleFunc("/a",handler_a)
   http.ListenAndServe(":8000",nil);
}
