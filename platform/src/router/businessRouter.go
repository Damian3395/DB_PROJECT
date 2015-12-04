package router

import(
	"fmt"
	"net/http"
	"log"
	db "./db"
	util "./db/util"
)

//UpdateBusinessGeneral
func UpdateBusinessGeneral(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update Business General Information\n")
}

//UpdateBusinessAddress
func UpdateBusinessAddress(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update Business Address Information\n")
}

//UpdateBusinessAll
func UpdateBusinessAll(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update All Business Information\n")
}

//CreateCoupon
func CreateCoupon(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Create Coupon\n")
}

//GetActiveCoupons
func GetActiveCoupons(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Active Coupons\n")
}

//GetExpiredCoupons
func GetExpiredCoupons(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Expired Coupons\n")
}
