package main

import(
  "fmt"
  "log"
  "net/http"
)

func main(){
  fileserver := http.FileServer(http.Dir("./src"))
  http.Handle("/", fileserver)
  http.HandleFunc("/form", formhandler)
  http.HandleFunc("/hello", hellohandler)

  fmt.Printf("Starting server at port 8080...\n")
  if err :=http.ListenAndServe(":8080",nil); err != nil{
    log.Fatal(err)
  }
}

func hellohandler(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/hello" {
    http.Error(w, "404 Page not found!", http.StatusNotFound)
    return
  }
  if r.Method != "GET"{
    http.Error(w, "Method now found", http.StatusNotFound)
    return
  }
  fmt.Fprintf(w,"Hello World!")
}

func formhandler(w http.ResponseWriter, r *http.Request){
  if err := r.ParseForm(); err != nil{
    fmt.Fprintf(w, "ParseForm() error: %v", err)
    return
  }
  fmt.Fprintf(w, "POST req success!\n")

  name := r.FormValue("name")
  email := r.FormValue("email")
  fmt.Fprintf(w, "Name: %s\n", name)
  fmt.Fprintf(w, "Email: %s\n", email)
}
