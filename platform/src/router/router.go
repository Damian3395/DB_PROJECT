package router

import(
	"io"
	"fmt"
	"net/http"
	db "./db"
	util "./db/util"	
)

func HandleLoginUser(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Login For User\n")
	var user util.Login_struct
	util.LoginToJSON(r, &user)
	status, err := db.LoginUser(user)
	
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling User Login Request\n")	
	io.WriteString(w, status)
}

func HandleLoginBusiness(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Login For Business\n")
	var user util.Login_struct
	util.LoginToJSON(r, &user)
	status, err := db.LoginBusiness(user)
	
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Business Login Request\n")
	io.WriteString(w, status)
}

func HandleRegisterUser(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Register For User\n")
	var user util.RegisterUser_struct
	util.RegisterUserToJSON(r, &user)
	status, err := db.RegisterUser(user)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling User Register Request\n")
	io.WriteString(w, status)
}

func HandleRegisterBusiness(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Register For Business\n")
	var business util.RegisterBusiness_struct
	util.RegisterBusinessToJSON(r, &business)
	status, err := db.RegisterBusiness(business)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Business Register Request\n")
	io.WriteString(w, status)
}

func HandleImages(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Image Requests\n");
  	filePath := "../res/img/" + r.URL.Query().Get("img")
  	fmt.Println("Serving Image File: ", filePath) 
  	http.ServeFile(w, r, filePath)
	fmt.Println("Done Handling Image Request")
}


