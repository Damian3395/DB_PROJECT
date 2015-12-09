package db

import
(
	"strings"
	"fmt"
	"github.com/jmcvetta/neoism"
	"encoding/json"
	util "./util"
)

func getMatches(res *[]struct{A string `json:"n.ID"`}, query util.QuerySearch_struct) (error){
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return err	
	}

	cq := neoism.CypherQuery{
    Statement: `
    	MATCH (n:Activity)
		WHERE n.NAME = {key_words} OR
				n.MAIN_CATEGORY = {main} AND
				n.SUB_CATEGORY = {sub} OR
				n.CAMPUS = {campus} OR
				n.MIN_AGE >= {min_age} AND 
				n.MAX_AGE <= {max_age} OR
				n.MIN_PEOPLE >= {min_people} AND
				n.MAX_PEOPLE <= {max_people}
		RETURN n.ID		
    `,
    Parameters: neoism.Props{"key_words": query.KEY_WORDS,
							 "main": query.MAIN_CATEGORY,
							 "sub": query.SUB_CATEGORY,
							 "campus": query.CAMPUS,
							 "min_age": query.MIN_AGE,
							 "max_age": query.MAX_AGE,
							 "min_people": query.MIN_PEOPLE,
							 "max_people": query.MAX_PEOPLE},
    Result:     res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return err
	}
		
	return nil
}

func getCouponsFromMatchs(business []struct{A string `json:"n.ID"`}) (string, int, error){
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "", 0, err	
	}

	coupon := []struct{
		A string `json:"n.COUPON_ID"` 
	}{}

	size := 0
	coupons := ""	
	for i := 0; i < len(business); i++ {
		id := business[i].A
		
		cq := neoism.CypherQuery{
		Statement: `
			MATCH (n:Coupon)
			WHERE n.ID = {user_id} AND n.VALID = "TRUE"
			RETURN n.COUPON_ID
		`,
		Parameters: neoism.Props{"user_id": id},
		Result:     &coupon,
		}
		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "", 0, err
		}

		if i > 0 && len(coupon) != 0 {
			coupons += ","
		}

		for x:=0; x < len(coupon); x++ {
			size++
			coupons += coupon[x].A
			if x != (len(coupon)-1) {
				coupons += ","
			}
		}
	}

	return coupons, size, nil
}

func getUserTickets(res *[]struct{A string `json:"n.COUPON_ID"`}, id string) (error){
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return err	
	}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (n:Ticket)
			WHERE n.ID = {user_id}
			RETURN n.COUPON_ID
		`,
		Parameters: neoism.Props{"user_id": id},
		Result:     &res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return err
	}

	return nil
}

func QueryCoupon(query util.QuerySearch_struct)	(string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering QueryCoupon\n")
	
	bestMatches := []struct{
		A string `json:"n.ID"`
	}{}	

	tickets := []struct{
		A string `json:"n.COUPON_ID"`
	}{}

	err := getMatches(&bestMatches, query)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	
	matchedCoupons, size, err := getCouponsFromMatchs(bestMatches)
	if err != nil {
		fmt.Printf("%s\n", err)	
	}

	err = getUserTickets(&tickets, query.ID)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	coupons := strings.Split(matchedCoupons,",")	

	for i:= 0; i < len(coupons); i++ {
		for x:= 0; x < len(tickets); x++ {
			if coupons[i] == tickets[x].A {
				coupons[i] = ""
				size--
				x = len(tickets)
			}
		}
	}

	if size == 0 {
		return "No Matches Found", nil
	}

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}

	response := ""
	data := map[string]interface{}{}
	jsonString := "{\"coupons\": ["
	for i:=0; i < len(coupons); i++ {
		id := coupons[i]
		res := []struct {
			N neoism.Node
		}{}

		if coupons[i] != ""{
			cq := neoism.CypherQuery{
				Statement: `
				MATCH (n:Coupon)
				WHERE n.COUPON_ID = {id}
				RETURN n
			`,
			Parameters: neoism.Props{"id": id},
			Result:     &res,
			}
			err = database.Cypher(&cq)
			if err != nil {
				fmt.Printf("Error Cyphering To Database\n")
				return "Internal Service Error", err
			}
		
			if len(res) != 0 {	
				couponStr, err := json.Marshal(res[0].N.Data)
				if err != nil {
					fmt.Printf("%s\n", err)
				}					
				jsonString += string(couponStr)
				if i != (len(coupons)-1){
					jsonString += ","
				}
			}
		}
	}	
	jsonString += "]}"

	dec := json.NewDecoder(strings.NewReader(jsonString))	
	dec.Decode(&data)
	couponSTR, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	response = string(couponSTR)

	return response, nil
}
