package main

import(
	"fmt"
	"net/http"
	router "./router"
)

func main(){
	fmt.Printf("Starting Server\n")
	http.Handle("/", http.FileServer(http.Dir("../../platform_ui/build")))
	http.Handle("/Application", http.FileServer(http.Dir("../../platform_ui/app")))
	http.HandleFunc("/LoginUser", router.HandleLoginUser)
	http.HandleFunc("/LoginBusiness", router.HandleLoginBusiness)
	http.HandleFunc("/RegisterUser", router.HandleRegisterUser)
	http.HandleFunc("/RegisterBusiness", router.HandleRegisterBusiness)
	http.HandleFunc("/res", router.HandleImages)
	http.ListenAndServe(":8080", nil)
}
