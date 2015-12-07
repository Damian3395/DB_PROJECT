package db

import
(
	"strings"
	"fmt"
	"encoding/json"
	"github.com/jmcvetta/neoism"
	util "./util"
)

func CreateCoupon(create util.CreateCoupon_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering CreateCoupon\n")
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
        A   string `json:"n.ID"` 
    }{}

	cq := neoism.CypherQuery{
    Statement: `
    	MATCH (n:Coupon)
      	WHERE n.ID = {user_id} AND n.VALID = "TRUE"
       	RETURN n.ID
    `,
    Parameters: neoism.Props{"user_id": create.ID},
    Result:     &res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if(len(res) < 5){
		id := GenerateCouponID()
		fmt.Printf("ID: %s", id)	
		success, err := FindCoupon(id)
		for (success == true) {
			id = GenerateCouponID()
			success, err = FindCoupon(id)
		}

		if err != nil {
			fmt.Printf("%s\n")
		}	

		cq = neoism.CypherQuery{
		Statement: `
			MATCH (a:Activity)
			WHERE a.ID = {user_id}
			CREATE (n:Coupon {ID: {user_id},
						NAME: a.NAME,
						ADDRESS: a.TOWNSHIP,
						CAMPUS: a.CAMPUS,
						TYPE: {type},
						DAY: {day},
						MONTH: {month},
						YEAR: {year},
						VALID: {valid},
						STUDENT: {student},
						COUPON_ID: {id}
				})
		`,
		Parameters: neoism.Props{"user_id": create.ID,
						"type": create.TYPE,
						"day": create.DAY,
						"month": create.MONTH,
						"year": create.YEAR,
						"valid": create.VALID,
						"student": create.STUDENT,
						"id": id},
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Internal Service Error", err
		}

		return "Success", nil	
	}else{
		fmt.Printf("Business Reached Max Coupons\n")
		return "Max Active Coupons Reached", nil
	}
}

func RemoveCoupon(coupon util.GetCoupon_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering RemoveCoupon\n")
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	cq := neoism.CypherQuery{
    Statement: `
    	MATCH (n:Coupon)
      	WHERE n.ID = {user_id} AND n.COUPON_ID = {coupon_id} 
       	DELETE n
    `,
    Parameters: neoism.Props{"user_id": coupon.ID, "coupon_id": coupon.COUPON_ID},
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	return "Success", nil
}

func GetActiveCoupons(coupon util.GetCoupon_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering GetActiveCoupon\n")
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
      	WHERE n.ID = {user_id} AND n.VALID = "TRUE"
       	RETURN n
    `,
    Parameters: neoism.Props{"user_id": coupon.ID},
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
		jsonString := "{\"coupons\": ["
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

func GetExpiredCoupons(coupon util.GetCoupon_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering GetExpiredCoupon\n")
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
      	WHERE n.ID = {user_id} AND n.VALID = "False"
       	RETURN n
    `,
    Parameters: neoism.Props{"user_id": coupon.ID},
    Result:     &res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if len(res) != 0 {
		response := ""
		coupon, err := json.Marshal(res[0].N.Data)
		response = string(coupon)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		return response, nil
	}

	return "No Expired Coupons", nil
}
