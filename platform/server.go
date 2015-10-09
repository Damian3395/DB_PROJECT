package main

import 
  (
    "fmt"
    "net/http"
  )

func handleLogin(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Login\n")
}

func handleAddDocument(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Add Document\n")
}

func handleUpdateDocument(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Update Document\n")
}

func handleRemoveDocument(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Remove Document\n")
}

func main() {
  fmt.Printf("Starting Server...\n")
  http.HandleFunc("/Login", handleLogin)
  http.HandleFunc("/AddDocument", handleAddDocument)
  http.HandleFunc("/UpdateDocument", handleUpdateDocument)
  http.HandleFunc("/RemoveDocument", handleRemoveDocument)
  http.ListenAndServe(":8080", http.FileServer(http.Dir("../platform_ui/build")))
}
