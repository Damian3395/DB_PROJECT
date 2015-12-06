package db

import
(
	"fmt"
	"github.com/jmcvetta/neoism"
	util "./util"
)

func UpdateUserGeneral(update util.UpdateGeneralUser_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering UpdateUserGeneral\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
		A	string `json:"n.FIRST_NAME"`
    }{}
	
	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Consumer)
       		WHERE n.USER_ID = {user_id}
			SET n.FIRST_NAME = {firstName}, n.LAST_NAME = {lastName}, n.AGE = {age}, n.GENDER = {gender}
        	RETURN n.FIRST_NAME
    	`,
    	Parameters: neoism.Props{"user_id": update.ID, "firstName": update.FIRST_NAME, 
									"lastName": update.LAST_NAME, "age": update.AGE, "gender": update.GENDER},
    	Result:     &res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if len(res) == 1 {
		return "Success", nil
	}

	return "Error: User Not Found", nil
}

func UpdateUserAddress(update util.UpdateAddressUser_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering UpdateUserAddress\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
		A	string `json:"n.ADDRESS"`
    }{}
	
	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Consumer)
       		WHERE n.USER_ID = {user_id}
        	SET n.ADDRESS = {address}, n.TOWNSHIP = {township}, n.STATE = {state}, n.ZIPCODE = {zipcode}
			RETURN n.ADDRESS
    	`,
    	Parameters: neoism.Props{"user_id": update.ID, "address": update.ADDRESS, "township": update.TOWNSHIP,
									"state": update.STATE, "zipcode": update.ZIPCODE},
    	Result:     &res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if len(res) == 1 {
		return "Success", nil
	}

	return "Error: User Not Found", nil
}

func UpdateUserStudent(update util.Student_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering UpdateUserStudent\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
		A	string `json:"n.RUID"`
    }{}
	
	cq := neoism.CypherQuery{
    	Statement: `
		 	MATCH (n:Student)
			WHERE n.USER_ID = {user_id}
			SET n.RUID = {ruid}, n.DEGREE = {degree}, n.CAMPUS = {campus}, n.YEAR = {year}
		 	RETURN n.ruid
		`,
		Parameters: neoism.Props{"user_id": update.ID, "ruid": update.RUID, "degree": update.DEGREE,
									"campus": update.CAMPUS, "year": update.YEAR},
		Result:     &res,
	}

	err = database.Cypher(&cq)		
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}
	
	if len(res) == 1 {
		return "Success", nil
	}else{
		cq := neoism.CypherQuery{
    	Statement: `
		 	CREATE (n:Student {USER_ID: {user_id}, RUID: {ruid}, DEGREE: {degree}, CAMPUS: {campus}, YEAR: {year}})
		`,
		Parameters: neoism.Props{"user_id": update.ID, "ruid": update.RUID, "degree": update.DEGREE,
									"campus": update.CAMPUS, "year": update.YEAR},
		Result:     &res,
	}

	err = database.Cypher(&cq)		
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}
	
		return "Success", nil
	}
}

func UpdateAllUser(update util.UpdateAllUser_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering UpdateUserStudent\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}

	res2 := []struct {
		A	string `json:"n.FIRST_NAME"`
    }{}


	res := []struct {
		A	string `json:"n.RUID"`
    }{}

	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Consumer)
       		WHERE n.USER_ID = {user_id}
			SET n.FIRST_NAME = {firstName}, n.LAST_NAME = {lastName}, n.AGE = {age}, n.GENDER = {gender},
									n.ADDRESS = {address}, n.TOWNSHIP = {township}, n.STATE = {state},
									n.ZIPCODE = {zipcode}
        	RETURN n.FIRST_NAME
    	`,
    	Parameters: neoism.Props{"user_id": update.ID, "firstName": update.FIRST_NAME, 
									"lastName": update.LAST_NAME, "age": update.AGE, "gender": update.GENDER,
									"address": update.ADDRESS, "township": update.TOWNSHIP, "state": update.STATE,
									"zipcode": update.ZIPCODE},
    	Result:     &res2,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return "Internal Service Error", err
	}

	if len(res2) != 1 {
		return "Failed", nil
	}
	
	if(update.RUID == "remove"){
		cq = neoism.CypherQuery{
	    	Statement: `
			 	MATCH (n:Student)
				WHERE n.USER_ID = {user_id}
				DELETE n
			`,
			Parameters: neoism.Props{"user_id": update.ID},
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Internal Service Error", err
		}

		return "Success", nil
	}else{
		cq = neoism.CypherQuery{
	    	Statement: `
			 	MATCH (n:Student)
				WHERE n.USER_ID = {user_id}
				SET n.RUID = {ruid}, n.DEGREE = {degree}, n.CAMPUS = {campus}, n.YEAR = {year}
			 	RETURN n.ruid
			`,
			Parameters: neoism.Props{"user_id": update.ID, "ruid": update.RUID, "degree": update.DEGREE,
										"campus": update.CAMPUS, "year": update.YEAR},
			Result:     &res,
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Internal Service Error", err
		}

		if len(res) == 1 {
			return "Success", nil
		}
	}

	return "Error: User Not Found", nil
}


