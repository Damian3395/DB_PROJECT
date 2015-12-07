package db

import
(
	"fmt"
	"strings"
	"encoding/json"
	"github.com/jmcvetta/neoism"
	util "./util"
)

func CreateTicket(tickets util.Ticket_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering CreateTicket\n")
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
    	MATCH (n:Ticket)
      	WHERE n.ID = {user_id} AND n.VALID = "TRUE"
       	RETURN n.ID
    `,
    Parameters: neoism.Props{"user_id": tickets.ID},
    Result:     &res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if(len(res) < 3){
		cq = neoism.CypherQuery{
		Statement: `
			CREATE (n:Ticket {ID: {user_id},
						COUPON_ID: {id},
						VALID: "TRUE"
				})
		`,
		Parameters: neoism.Props{"user_id": tickets.ID,
						"id": tickets.COUPON_ID},
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Internal Service Error", err
		}

		return "Success", nil	
	}else{
		fmt.Printf("User Reached Max Tickets\n")
		return "Max Active Tickets Reached", nil
	}
}

func UseTicket(tickets util.Ticket_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering UseTicket\n")
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	cq := neoism.CypherQuery{
    Statement: `
		MATCH (t:Ticket)
		WHERE t.ID = {user_id} AND t.COUPON_ID = {id}
		SET t.VALID = "FALSE"
	`,
    Parameters: neoism.Props{"user_id": tickets.ID, "id": tickets.COUPON_ID},
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}
		
	return "Success", nil
}

func GetUserValidTickets(user util.User_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering GetUserValidTickets\n")
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}

	res := []struct {
		A   string `json:"n.COUPON_ID"`
	}{}

	res2 := []struct {
		N neoism.Node
	}{}
	
	cq := neoism.CypherQuery{
    Statement: `
		MATCH (n:Ticket)
		WHERE n.VALID = "TRUE" AND n.ID = {user_id}
		RETURN n.COUPON_ID
	`,
    Parameters: neoism.Props{"user_id": user.ID},
	Result:		&res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if len(res) == 0 {
		return "Active Coupons Does Not Exist", nil
	}else{
		fmt.Printf("%s\n", res[0].A)
	}

	response := ""
	data := map[string]interface{}{}
	jsonString := "{\"coupons\": ["
	for i := 0; i < len(res); i++ {
		cq := neoism.CypherQuery{
		Statement: `
			MATCH (n:Coupon)
			WHERE n.COUPON_ID = {id}
			RETURN n
		`,
		Parameters: neoism.Props{"id": res[i].A},
		Result:		&res2,
		}
		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Internal Service Error", err
		}
		
		if len(res2) != 0 {
			coupon, err := json.Marshal(res2[0].N.Data)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			jsonString += string(coupon)
			if i != (len(res)-1){
				jsonString += ","
			}
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

func GetUserInvalidTickets(user util.User_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server Cyphering GetUseriInValidTickets\n")
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
      	WHERE n.ID = {user_id} AND n.VALID = "False"
		RETURN n
	`,
    Parameters: neoism.Props{"user_id": user.ID},
	Result:		&res,
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
	
	return "Used Coupon Does Not Exist", nil
}
