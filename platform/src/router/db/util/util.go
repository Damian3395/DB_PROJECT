package util

import(
	"log"
	"encoding/json"
	"net/http"
)

const(
	TRUE = 1
	FALSE = 0
	ERROR_
)

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

func PrintLogin(login *Login_struct){
	obj, err := json.Marshal(login)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(obj))
}

func PrintRegisterUser(register *RegisterUser_struct){
	obj, err := json.Marshal(register)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(obj))
}

func PrintRegisterBusiness(register *RegisterBusiness_struct){
	obj, err := json.Marshal(register)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(obj))
}

func PrintConsumer(consumer *Consumer_struct){
	obj, err := json.Marshal(consumer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(obj))
}

func PrintBusiness(business *Business_struct){
	obj, err := json.Marshal(business)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(obj))
}

func PrintStudent(student *Student_struct){
	obj, err := json.Marshal(student)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(obj))
}


