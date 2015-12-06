package main

import(
	"fmt"
	"net/http"
	router "./router"
)

func main(){
	fmt.Printf("Starting Server\n")
	//Serve Webpage
	http.Handle("/", http.FileServer(http.Dir("../../platform_ui/build")))
	
	//Serve Basic Webpage Functions
	http.HandleFunc("/LoginUser", router.HandleLoginUser)
	http.HandleFunc("/LoginBusiness", router.HandleLoginBusiness)
	http.HandleFunc("/RegisterUser", router.HandleRegisterUser)
	http.HandleFunc("/RegisterBusiness", router.HandleRegisterBusiness)
	http.HandleFunc("/res", router.HandleImages)

	//Serve User Application Functions
	http.HandleFunc("/GetUserInformation", router.GetUserInformation)
	http.HandleFunc("/GetStudentInformation", router.GetStudentInformation)
	http.HandleFunc("/UpdateUserGeneral", router.UpdateUserGeneral)
	http.HandleFunc("/UpdateUserAddress", router.UpdateUserAddress)
	http.HandleFunc("/UpdateUserStudent", router.UpdateUserStudent)
	http.HandleFunc("/UpdateUserAll", router.UpdateUserAll)
	http.HandleFunc("/GetUsedCoupons", router.GetUsedCoupons)
	http.HandleFunc("/GetCurrentCoupons", router.GetCurrentCoupons)
	http.HandleFunc("/QueryCoupon", router.QueryCoupon)
	http.HandleFunc("/OptimizeCoupon", router.OptimizeCoupon)

	//Serve Business Application Functions
	http.HandleFunc("/GetBusinessInformation", router.GetBusinessInformation)
    http.HandleFunc("/UpdateBusinessGeneral", router.UpdateBusinessGeneral)
	http.HandleFunc("/UpdateBusinessAddress", router.UpdateBusinessAddress)
	http.HandleFunc("/UpdateBusinessAll", router.UpdateBusinessAll)
	http.HandleFunc("/CreateCoupon", router.CreateCoupon)
	http.HandleFunc("/RemoveCoupon", router.RemoveCoupon)
	http.HandleFunc("/GetActiveCoupons", router.GetActiveCoupons)
	http.HandleFunc("/GetExpiredCoupons", router.GetExpiredCoupons)	

	//Listen On Port 8080
	http.ListenAndServe(":8080", nil)
}
