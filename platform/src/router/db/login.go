package db

import
(
	"fmt"
	"github.com/jmcvetta/neoism"
	Util "./util"
)

func LoginUser(login Util.Login_struct) (string, error) {
	fmt.Printf("Accessing NEO4J Server Cyphering UserLogin\n")
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
        A   string `json:"n.USER_ID"` 
    }{}

	cq := neoism.CypherQuery{
    Statement: `
    	MATCH (n:User)
      	WHERE n.USER_ID = {user_id} AND n.PASSWORD = {password}
       	RETURN n.USER_ID
    `,
    Parameters: neoism.Props{"user_id": login.ID, "password": login.PASSWORD},
    Result:     &res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if(len(res) == 1){
		fmt.Printf("Successful Login\n")
		return "Success", nil	
	}

	fmt.Printf("User Login Failed\n")
	return "Failed", nil
}

func LoginBusiness(login Util.Login_struct) (string, error) {
	fmt.Printf("Accessing NEO4J Server Cyphering BusinessLogin\n")
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}

	res := []struct {
        	A   string `json:"n.USER_ID"` 
    }{}
	
	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Business)
       		WHERE n.USER_ID = {user_id} AND n.PASSWORD = {password}
        	RETURN n.USER_ID
    	`,
    	Parameters: neoism.Props{"user_id": login.ID, "password": login.PASSWORD},
    	Result:     &res,
	}
	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if(len(res) == 1){
		return "Success", nil
	}

	fmt.Printf("Business User Login Failed\n")
	return "Failed", nil
}
