package router

import(
	"fmt"
	"net/http"
	"log"
	db "./db"
	util "./db/util"
)

//UpdateUserGeneral
func UpdateUserGeneral(w http.ResonseWriter, r *http.Request){
	fmt.Printf("Handling Update User General Information\n")
	
	//Do Stuff
	
	fmt.Printf("Done Handling Update User General Information\n")
}

//UpdateUserAddress
func UpdateUserAddress(w http.ResonseWriter, r *http.Request){
	fmt.Printf("Handling Update User Address Information\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Update User Address Information\n")
}

//UpdateUserStudent
func UpdateUserStudent(w http.ResonseWriter, r *http.Request){
	fmt.Printf("Handling Update User Student Information\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Update User Student Information\n")
}

//UpdateUserAll
func UpdateUserAll(w http.ResonseWriter, r *http.Request){
	fmt.Printf("Handling Update All User Information\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Update All User  Information\n")
}

//GetUsedCoupons
func GetUsedCoupons(w http.ResonseWriter, r *http.Request){
	fmt.Printf("Handling Get Used Coupons\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Get Used Coupons\n")
}

//GetCurrentCoupons
func GetCurrentCoupons(w http.ResonseWriter, r *http.Request){
	fmt.Printf("Handling Get Current Coupons\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Get Current Coupons\n")
}

//QueryCoupon
func QueryCoupon(w http.ResonseWriter, r *http.Request){
	fmt.Printf("Handling Query Coupon\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Query Coupon\n")
}

//OptimizeCoupon
func OptimizeCoupon(w http.ResonseWriter, r *http.Request){
	fmt.Printf("Handling Optimize Coupon\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Optimize Coupon\n")
}


