package db

import
(
	"fmt"
	"strings"
	"github.com/jmcvetta/neoism"
	"encoding/json"
	util "./util"
)

func QueryCoupon(query util.QuerySearch_struct)	(string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering QueryCoupon\n")
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
		
	bestMatch := []struct {
		A string `json:"n.ID"`		 
    }{}

	used := []struct {
		A string `json:"n.COUPON_ID"`
	}{}
	
	
	cq := neoism.CypherQuery{
    Statement: `
    	MATCH (n:Activity)
		WHERE n.NAME = {key_words} OR
				n.MAIN_CATEGORY = {main} OR
				n.SUB_CATEGORY = {sub} OR
				n.CAMPUS = {campus} OR
				n.MIN_AGE >= {min_age} AND
				n.MAX_AGE <= {max_age} OR
				n.MIN_PEOPLE >= {min_people} and
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
    Result:     &bestMatch,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

		fmt.Printf("%s\n", len(bestMatch))
		cq = neoism.CypherQuery{
		Statement: `
			MATCH (t:Ticket)
			WHERE t.ID = "user_id"
			RETURN t.COUPON_ID
		`,
		Parameters: neoism.Props{"user_id": query.ID},
		Result:     &used,
		}
		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Internal Service Error", err
		}


	fmt.Printf("%s\n", len(used))
	//response := ""
	//data := map[string]interface{}{}
	//jsonString := "{\"coupons\": ["
	for i := 0; i < len(bestMatch); i++ {
		coupons := []struct {
			N neoism.Node 
		}{}
		cq := neoism.CypherQuery{
		Statement: `
			MATCH (c:Coupon)
			WHERE c.ID = {user_id} AND c.VALID = "TRUE"
			RETURN c 
		`,
		Parameters: neoism.Props{"user_id": bestMatch[i].A},
		Result:     &coupons,
		}
		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Internal Service Error", err
		}
		
		
		for i := 0; i < len(coupons); i++ {
			coupon, err := json.Marshal(coupons[i].N.Data)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			fmt.Printf("%s\n", string(coupon))
		}
	}
/*
	response := ""
	data := map[string]interface{}{}
	if len(res) != 0 {
		jsonString := "{\"objects\": ["
		for i:=0; i < len(res); i++ {
			coupon, err := json.Marshal(res[i].N.Data)
			if err != nil {
				fmt.Printf("%s\n", err)
			}	
			jsonString += string(coupon)
			if i != (len(res)-1){
				jsonString += ","
			}
		}
		jsonString += "]}"
		
		dec := json.NewDecoder(strings.NewReader(jsonString))
		dec.Decode(&data)
		
		coupons, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		response = string(coupons)
			
		return response, nil
	}
*/

	return "", nil
}

func OptimizedSearch(query util.User_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering OptimizedSearch\n")
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
		N neoism.Node 
    }{}

	cq := neoism.CypherQuery{
    Statement: `
    	MATCH (n:Ticket)
      	WHERE n.USER_ID = {user_id}
       	RETURN n
    `,
    Parameters: neoism.Props{"user_id": query.ID},
    Result:     &res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}


	response := ""
	data := map[string]interface{}{}
	if len(res) != 0 {
		jsonString := "{\"objects\": ["
		for i:=0; i < len(res); i++ {
			coupon, err := json.Marshal(res[i].N.Data)
			if err != nil {
				fmt.Printf("%s\n", err)
			}	
			jsonString += string(coupon)
			if i != (len(res)-1){
				jsonString += ","
			}
		}
		jsonString += "]}"
		
		dec := json.NewDecoder(strings.NewReader(jsonString))
		dec.Decode(&data)
		
		coupons, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		response = string(coupons)
			
		return response, nil
	}

	return "No Active Coupons", nil
}
