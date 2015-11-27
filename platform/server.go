package main

import 
  (
	"encoding/json"
	"log"
	"io"
    "fmt"
    "net/http"
  )

type login_struct struct {
	ID string
	PASSWORD string
}

type registerUser_struct struct {
	FIRST_NAME string
	LAST_NAME string
	AGE string
	STUDENT string
	ADDRESS string
	TOWNSHIP string
	STATE string
	ZIPCODE string
	RUID string
	DEGREE string
	CAMPUS string
	YEAR string
}

type registerBusiness_struct struct {
	NAME string
	ADDRESS string
	TOWNSHIP string
	MIN_AGE string
	MAX_AGE string
	MIN_PEOPLE string
	MAX_PEOPLE string
	MAIN_CATEGORY string
	SUB_CATEGORY string
}

func handleLoginUser(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Login For User\n")
  
  decoder := json.NewDecoder(r.Body)
  var login login_struct
  err := decoder.Decode(&login)
  if err != nil {
	log.Fatal(err)
  }	

  io.WriteString(w, "Logging In User Through Proxy Server")
}

func handleLoginBusiness(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Login For Business\n")

  decoder := json.NewDecoder(r.Body)
  var login login_struct
  err := decoder.Decode(&login)
  if err != nil {
    log.Fatal(err)
  }

  log.Println("User: ", login.ID)
  log.Println("Password: ", login.PASSWORD)

  io.WriteString(w, "Logging In Business Through Proxy Server")
}

func handleRegisterUser(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Register For User\n")

  decoder := json.NewDecoder(r.Body)
  var register registerUser_struct 
  err := decoder.Decode(&register)
  if err != nil {
    log.Fatal(err)
  }

  jsonObj, err := json.Marshal(&register)
  if err != nil {
    log.Fatal(err)
  }
  log.Println(string(jsonObj))

  io.WriteString(w, "Registering New User Through Proxy Server")
}

func handleRegisterBusiness(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Register For Business\n")
  
  decoder := json.NewDecoder(r.Body)
  var register registerBusiness_struct
  err := decoder.Decode(&register)
  if err != nil {
    log.Fatal(err)
  }

  jsonObj, err := json.Marshal(&register)
  if err != nil {
    log.Fatal(err)
  }
  log.Println(string(jsonObj))

  io.WriteString(w, "Registering New Business Through Proxy Server")
}

func handleImages(w http.ResponseWriter, r *http.Request){
  fmt.Printf("Handling Image Requests\n");
 
  filePath := "res/img/" + r.URL.Query().Get("img")
  fmt.Println("Image File Path: ", filePath) 
  http.ServeFile(w, r, filePath)
}

func main() {
  fmt.Printf("Starting Server...\n")
  http.Handle("/", http.FileServer(http.Dir("../platform_ui/build")))
  http.HandleFunc("/LoginUser", handleLoginUser)
  http.HandleFunc("/LoginBusiness", handleLoginBusiness)
  http.HandleFunc("/RegisterUser", handleRegisterUser)
  http.HandleFunc("/RegisterBusiness", handleRegisterBusiness)
  http.HandleFunc("/res", handleImages)
  http.ListenAndServe(":8080", nil)
}
