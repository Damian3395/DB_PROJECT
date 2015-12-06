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
	
	res := []struct {
		N neoism.Node 
    }{}

	cq := neoism.CypherQuery{
    Statement: `
    	MATCH (n:Coupon)
      	WHERE 
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
