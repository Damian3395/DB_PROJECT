package db

import
(
	"fmt"
	"encoding/json"
	"github.com/jmcvetta/neoism"
)

func GetUserInfo(userID string) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering GetUserInfo %s\n", userID)

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
        	MATCH (n:Consumer)
       		WHERE n.USER_ID = {user_id}
        	RETURN n
    	`,
    	Parameters: neoism.Props{"user_id": userID},
    	Result:     &res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if len(res) == 1 {
		response := ""
		consumer, err := json.Marshal(res[0].N.Data)
		response = string(consumer)
		if(err != nil){
			fmt.Printf("%s\n", err)
		}
		
		return response, nil
	}

	return "Error: User Not Found", nil
}

func GetStudentInfo(userID string) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering GetStudentInfo %s\n", userID)

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
        	MATCH (n:Student)
       		WHERE n.USER_ID = {user_id}
        	RETURN n
    	`,
    	Parameters: neoism.Props{"user_id": userID},
    	Result:     &res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if len(res) == 1 {
		response := ""
		student, err := json.Marshal(res[0].N.Data)
		response = string(student)
		if(err != nil){
			fmt.Printf("%s\n", err)
		}
		
		return response, nil
	}

	return "Error: User Not Found", nil
}
