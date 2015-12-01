package router

import(
	"fmt"
	"net/http"
	"log"
	db "./db"
	util "./db/util"	
)

func HandleLoginUser(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Login For User\n")
	var login util.Login_struct
	util.LoginToJSON(r, &login)
	loginSuccess, err := db.LoginUser(login)
	if err != nil {
		log.Fatal(err)	
	}

	if loginSuccess {
		fmt.Printf("Done Handling User Login Request\n")
	}
}

func HandleLoginBusiness(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Login For Business\n")
	var login util.Login_struct
	util.LoginToJSON(r, &login)
	loginSuccess, err := db.LoginBusiness(login)
	if err != nil {
		log.Fatal(err)	
	}

	if loginSuccess {
		fmt.Printf("Done Handling Business Login Request\n")
	}
}

func HandleRegisterUser(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Register For User\n")
	var user util.RegisterUser_struct
	util.RegisterUserToJSON(r, &user)
	registerSuccess, err := db.RegisterUser(user)
	if err != nil {
		log.Fatal(err)	
	}

	if registerSuccess {
		fmt.Printf("Done Handling User Register Request\n")
	}
}

func HandleRegisterBusiness(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Register For Business\n")
	var business util.RegisterBusiness_struct
	util.RegisterBusinessToJSON(r, &business)
	registerSuccess, err := db.RegisterBusiness(business)
	if err != nil {
		log.Fatal(err)	
	}

	if registerSuccess {
		fmt.Printf("Done Handling Business Register Request\n")
	}
}

func HandleImages(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Image Requests\n");
  	filePath := "../res/img/" + r.URL.Query().Get("img")
  	fmt.Println("Serving Image File: ", filePath) 
  	http.ServeFile(w, r, filePath)
	fmt.Println("Done Handling Image Request")
}


