package router

import(
	"io"
	"fmt"
	"net/http"
	"io/ioutil"
	db "./db"
	util "./db/util"
)

//GetBusinessInformation
func GetBusinessInformation(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Business Information\n")
	
	var user util.User_struct
	util.UserToJSON(r, &user)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.GetBusinessInfo(user.ID)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Get Business Information\n")
	io.WriteString(w, status)
}


//UpdateBusinessGeneral
func UpdateBusinessGeneral(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update Business General Information\n")

	var update util.UpdateGeneralBusiness_struct
	util.GeneralBusinessToJSON(r, &update)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.UpdateBusinessGeneral(update)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Update Business General Information\n")
	io.WriteString(w, status)
}

//UpdateBusinessAddress
func UpdateBusinessAddress(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update Business Address Information\n")

	var update util.UpdateAddressBusiness_struct
	util.AddressBusinessToJSON(r, &update)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.UpdateBusinessAddress(update)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Update Business Address Information\n")
	io.WriteString(w, status)
}

//UpdateBusinessAll
func UpdateBusinessAll(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update All Business Information\n")

	var update util.Business_struct
	util.BusinessToJSON(r, &update)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.UpdateBusinessAll(update)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	
	fmt.Printf("Done Handling Update All Business Information")
	io.WriteString(w, status)
}

//CreateCoupon
func CreateCoupon(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Create Coupon\n")
	
	var coupon util.CreateCoupon_struct
	util.CreateCouponToJSON(r, &coupon)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.CreateCoupon(coupon)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Create Coupon\n")
	io.WriteString(w, status)
}

func RemoveCoupon(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Remove Coupon\n")
	
	var coupon util.GetCoupon_struct
	util.GetCouponToJSON(r, &coupon)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.RemoveCoupon(coupon)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	
	fmt.Printf("Done Handling Remove Coupon\n")
	io.WriteString(w, status)
}

//GetActiveCoupons
func GetActiveCoupons(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Active Coupons\n")

	var coupon util.GetCoupon_struct
	util.GetCouponToJSON(r, &coupon)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.GetActiveCoupons(coupon)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	
	fmt.Printf("Done Handling Get Active Coupons\n")
	io.WriteString(w, status)
}

//GetExpiredCoupons
func GetExpiredCoupons(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Expired Coupons\n")

	var coupon util.GetCoupon_struct
	util.GetCouponToJSON(r, &coupon)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.GetExpiredCoupons(coupon)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Get Expired Coupons\n")
	io.WriteString(w, status)
}

//GetCouponAnalytics
func GetCouponAnalytics(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Coupon Analytics\n")
	
	var analyze util.Analyze_struct
	util.AnalyzeToJSON(r, &analyze)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.AnalyzeCoupon(analyze)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Get Expired Coupons\n")
	io.WriteString(w, status)
}
