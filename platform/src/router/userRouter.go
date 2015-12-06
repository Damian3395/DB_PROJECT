package router

import(
	"io"
	"fmt"
	"net/http"
	"io/ioutil"
	db "./db"
	util "./db/util"
)

//GetUserInformation
func GetUserInformation(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get User Information\n")
	
	var user util.User_struct
	util.UserToJSON(r, &user)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.GetUserInfo(user.ID)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Get Use Information\n")
	io.WriteString(w, status)
}

func GetStudentInformation(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Student Information\n")
	
	var user util.User_struct
	util.UserToJSON(r, &user)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))

	status, err := db.GetStudentInfo(user.ID)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Get Student Information\n")
	io.WriteString(w, status)

}

//UpdateUserGeneral
func UpdateUserGeneral(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update User General Information\n")
	
	var update util.UpdateGeneralUser_struct
	util.GeneralUserToJSON(r, &update)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.UpdateUserGeneral(update)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Update User General Information\n")
	io.WriteString(w, status)
}

//UpdateUserAddress
func UpdateUserAddress(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update User Address Information\n")

	var update util.UpdateAddressUser_struct
	util.AddressUserToJSON(r, &update)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.UpdateUserAddress(update)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Update User Address Information\n")
	io.WriteString(w, status)
}

//UpdateUserStudent
func UpdateUserStudent(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update User Student Information\n")

	var update util.Student_struct
	util.StudentToJSON(r, &update)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.UpdateUserStudent(update)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Update User Student Information\n")
	io.WriteString(w, status)
}

//UpdateUserAll
func UpdateUserAll(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Update All User Information\n")

	var update util.UpdateAllUser_struct
	util.UserAllToJSON(r, &update)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.UpdateAllUser(update)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Update All User  Information\n")
	io.WriteString(w, status)
}

//GetUsedCoupons
func GetUsedCoupons(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Used Coupons\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Get Used Coupons\n")
	io.WriteString(w, "User Used Coupons Retrieved\n")
}

//GetCurrentCoupons
func GetCurrentCoupons(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get Current Coupons\n")

	//Do Stuff
	
	fmt.Printf("Done Handling Get Current Coupons\n")
	io.WriteString(w, "Use Active Coupons Retrieved\n")
}

//QueryCoupon
func QueryCoupon(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Query Coupon\n")

	var query util.QuerySearch_struct
	util.QuerySearchToJSON(r, &query)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.QueryCoupon(query)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	
	fmt.Printf("Done Handling Query Coupon\n")
	io.WriteString(w, status)
}

//OptimizeCoupon
func OptimizeCoupon(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Optimize Coupon\n")

	var query util.User_struct
	util.UserToJSON(r, &query)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.OptimizedSearch(query)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	
	fmt.Printf("Done Handling Optimize Coupon\n")
	io.WriteString(w, status)
}



