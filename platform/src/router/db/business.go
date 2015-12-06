package db

import
(
	"fmt"
	"encoding/json"
	"github.com/jmcvetta/neoism"
)

func GetBusinessInfo(userID string) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering GetBusinessInfo %s\n", userID)

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
        	MATCH (n:Activity)
       		WHERE n.ID = {user_id}
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
		business, err := json.Marshal(res[0].N.Data)
		response = string(business)
		if(err != nil){
			fmt.Printf("%s\n", err)
		}
		
		return response, nil
	}

	return "Error: User Not Found", nil
}
