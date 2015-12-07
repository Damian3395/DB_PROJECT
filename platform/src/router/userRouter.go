package router

import(
	"io"
	"fmt"
	"net/http"
	"io/ioutil"
	db "./db"
	util "./db/util"
)

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

func GetUserInvalidTickets(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get User Invalid Tickets\n")

	var user util.User_struct
	util.UserToJSON(r, &user)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.GetUserInvalidTickets(user)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Done Handling Get User Invalid Tickets\n")
	io.WriteString(w, status)
}

func GetUserValidTickets(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Get User Valid Tickets\n")

	var user util.User_struct
	util.UserToJSON(r, &user)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.GetUserValidTickets(user)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	
	fmt.Printf("Done Handling Get User Valid Tickets\n")
	io.WriteString(w, status)
}

func UseTicket(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Use Ticket\n")

	var ticket util.Ticket_struct
	util.TicketToJSON(r, &ticket)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.UseTicket(ticket)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
		
	fmt.Printf("Done Handling Use Ticket\n")
	io.WriteString(w, status)
}

func CreateTicket(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Handling Create Ticket\n")

	var ticket util.Ticket_struct
	util.TicketToJSON(r, &ticket)
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", string(contents))
	
	status, err := db.CreateTicket(ticket)
	if err != nil {
		fmt.Printf("%s\n", err)
	}	
	fmt.Printf("Done Handling Create Ticket\n")
	io.WriteString(w, status)
}

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



