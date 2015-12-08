package db

import(
	"fmt"
	"github.com/jmcvetta/neoism"
	util "./util"
	"strconv"
	"strings"
)

func AnalyzeCoupon(analyze util.Analyze_struct)(string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering Analyze Coupon\n")	

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")	

	users := []struct {
		A string `json:"n.USER_ID"`
		B string `jsonL"n.VALID"`
	}{}

	cq := neoism.CypherQuery{
		Statement:`
			MATCH (n:Ticket)
			WHERE n.COUPON_ID = {id}
			RETURN n.USER_ID, n.VALID
		`,
		Parameters: neoism.Props{"id": analyze.COUPON_ID},
		Result:		&users,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Interna Service Error", err
	}

	if len(users) == 0 {
		return "None", nil
	}

	valid := 0
	invalid := 0
	male := 0
	female := 0
	grp1 := 0
	grp2 := 0
	grp3 := 0
	grp4 := 0
	grp5 := 0
	grp6 := 0
	inState := 0
	outState := 0
	for i := 0; i < len(users); i++{
		if(users[i].B == "TRUE"){
			valid++
		}else{
			invalid++
		}

		res := []struct {
			A string `json:"n.GENDER"`
			B string `json:"n.AGE"`
			C string `json:"n.STATE"`
		}{}

		cq := neoism.CypherQuery{
		Statement:`
			MATCH (n:Consumer)
			WHERE n.USER_ID = {id}
			RETURN n
		`,
		Parameters: neoism.Props{"id": users[i].A},
		Result:		&res,
		}
		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Interna Service Error", err
		}

		if res[0].A == "MALE" {
			male++
		}else{
			female++
		}
	
		age, err := strconv.Atoi(res[0].B)
		if err != nil {
			fmt.Printf("%s\n", err)
		}		

		if age >= 1 && age <= 13 {
			grp1++
		}else if age >= 14 && age <= 18 {
			grp2++
		}else if age >= 19 && age <= 21 {
			grp3++
		}else if age >= 22 && age <= 30 {
			grp4++
		}else if age >= 31 && age <= 50 {
			grp5++
		}else{
			grp6++
		}

		state := strings.ToLower(res[0].C)
		if state == "nj" || state == "new jersey"{
			inState++
		}else{
			outState++
		}
	}

	student := 0
	consumer := 0
	for i:=0; i < len(users); i++ {
		res := []struct {
			A string `json:"n.DEGREE"`
		}{}

		cq := neoism.CypherQuery{
		Statement:`
			MATCH (n:Student)
			WHERE n.USER_ID = {id}
			RETURN n.DEGREE
		`,
		Parameters: neoism.Props{"id": users[i].A},
		Result:		&res,
		}
		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Interna Service Error", err
		}

		if len(res) != 0 {
			student++
		}else{
			consumer++
		}
	}

	response := "{ valid: " + strconv.Itoa(valid) + 
					", invalid: " + strconv.Itoa(invalid) +
					", male: " + strconv.Itoa(male) +
					", female: " + strconv.Itoa(female) +
					", student: " + strconv.Itoa(student) +
					", consumer: " + strconv.Itoa(consumer) +
					", grp1: " + strconv.Itoa(grp1) +
					", grp2: " + strconv.Itoa(grp2) +
					", grp3: " + strconv.Itoa(grp3) +
					", grp4: " + strconv.Itoa(grp4) +
					", grp5: " + strconv.Itoa(grp5) +
					", grp6: " + strconv.Itoa(grp6) + "}" 

	return response, nil
}
