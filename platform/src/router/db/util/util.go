package util

import(
	"log"
	"encoding/json"
	"net/http"
)

func UserToJSON(r *http.Request, structure *User_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func TicketToJSON(r *http.Request, structure *Ticket_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func QuerySearchToJSON(r *http.Request, structure *QuerySearch_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateCouponToJSON(r *http.Request, structure *CreateCoupon_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCouponToJSON(r *http.Request, structure *GetCoupon_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}
func UserAllToJSON(r *http.Request, structure *UpdateAllUser_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func GeneralUserToJSON(r *http.Request, structure *UpdateGeneralUser_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func AddressUserToJSON(r *http.Request, structure *UpdateAddressUser_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func GeneralBusinessToJSON(r *http.Request, structure *UpdateGeneralBusiness_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func AddressBusinessToJSON(r *http.Request, structure *UpdateAddressBusiness_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}


func LoginToJSON(r *http.Request, structure *Login_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterUserToJSON(r *http.Request, structure *RegisterUser_struct) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterBusinessToJSON(r *http.Request, structure *RegisterBusiness_struct) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumerToJSON(r *http.Request, structure *Consumer_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func BusinessToJSON(r *http.Request, structure *Business_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}

func StudentToJSON(r *http.Request, structure *Student_struct){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Fatal(err)
	}
}
