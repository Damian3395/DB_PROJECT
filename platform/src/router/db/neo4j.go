package db

import
(
	"fmt"
	"github.com/jmcvetta/neoism"
	Util "./util"
)

/*
func FindConsumer(first_name string, last_name string, address string, state string, zipcode string) (bool, err){
	fmt.Printf("Accessing NEO4J Server: Cyphering Find Consumer\n")
	
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")

	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err
	}

	res := []struct {
		A	string 'json:"n.FIRST_NAME"'
	}{}

	cq := neoism.CypherQuery{
		Statement:'
			MATCH (n:Consumer)
			WHERE n.FIRST_NAME = {first_name} AND n.LAST_NAME = {last_name} AND n.ADDRESS = {address} AND n.STATE = {state} AND n.ZIPCODE = {zipcode}
			RETURN n.FIRST_NAME
		',
		Parameters: neoism.Props{"first_name": first_name, "last_name": last_name, "address": address, "state": state, "zipcode": zipcode},
		Result:		&res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return false, err
	}

	if len(res) == 1 {
		return true, nil
	}

	return false, nil
}

func FindStudent(ruid string) (bool, error) {
	fmt.Printf("Accessing NEO4J Server: Cyphering Find Student %s\n", ruid)

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")

	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err
	}

	res := []struct {
		A	string 'json:"n.RUID"'	
	}{}

	cq := neoism.CypherQuery{
		Statement:'
			MATCH (n:Student)
			WHERE n.RUID = {ruid}
			RETURN n.RUID
		',
		Parameters: neoism.Props{"ruid": ruid},
		Result:		&res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return false, err
	}

	if len(res) == 1 {
		return true, nil
	}

	return false, nil
}
*/

func FindUser(userID string) (bool, error) {
	fmt.Printf("Accessing NEO4J Server: Cyphering Find User %s\n", userID)

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err	
	}
	
	res := []struct {
        A   string `json:"n.USER_ID"` 
    }{}

	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:User)
       		WHERE n.USER_ID = {user_id}
        	RETURN n.USER_ID
    	`,
    	Parameters: neoism.Props{"user_id": userID},
    	Result:     &res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return false, err
	}

	if len(res) == 1 {
		return true, nil
	}

	return false, nil
}

func FindBusiness(userID string) (bool, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering Find Business %s\n", userID)
	
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err	
	}

	res := []struct {
        	A   string `json:"n.USER_ID"` 
    	}{}

	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Business)
       		WHERE n.USER_ID = {user_id}
        	RETURN n.USER_ID
    	`,
    	Parameters: neoism.Props{"user_id": userID},
    	Result:     &res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return false, err	
	}

	if len(res) == 1 {
		return true, nil
	}

	return false, nil
}

func LoginUser(login Util.Login_struct) (bool, error) {
	fmt.Printf("Accessing NEO4J Server Cyphering UserLogin\n")
	Util.PrintLogin(&login);

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err	
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
		return false, err
	}

	if len(res) == 1 {
		fmt.Printf("Successfully Logged In User\n")
		return true, nil
	}

	fmt.Printf("User Login Failed\n")
	return false, nil
}

func LoginBusiness(login Util.Login_struct) (bool, error) {
	fmt.Printf("Accessing NEO4J Server Cyphering BusinessLogin\n")
	Util.PrintLogin(&login);

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err	
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
		return false, err
	}

	if len(res) == 1 {
		fmt.Printf("Successfully Logged In Business User\n")
		return true, nil
	}

	fmt.Printf("Business User Login Failed\n")
	return false, nil
}

func RegisterUser(register Util.RegisterUser_struct) (bool, error) {
	fmt.Printf("Accessing NEO4J Server Cyphering User Register\n")
	Util.PrintRegisterUser(&register);

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err	
	}

	// Check If User Login Node Exists
	userExists, err := FindUser(register.USER_ID)
	if err != nil {
		return false, err
	}

	// Check If Student Node Exists
	//studentExists, err := FindStudent(register.RUID)
	//if err != nil {
	//	return false, err
	//}
	
	// Create Login Node
	if(!userExists){
		fmt.Printf("Creating User Node\n")
		cq := neoism.CypherQuery{
		Statement: `
			CREATE (n:User {USER_ID: {user_id}, PASSWORD: {password}})
		`,
		Parameters: neoism.Props{"user_id": register.USER_ID, "password": register.PASSWORD},
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return false, err
		}

		// Create Consumer Node
		fmt.Printf("Creating Consumer Node\n")
		cq = neoism.CypherQuery{
		Statement: `
			CREATE (n:Consumer {USER_ID: {user_id},
						FIRST_NAME: {first_name},
						LAST_NAME: {last_name},
						AGE: {age},
						ADDRESS: {address},
						TOWNSHIP: {township},
						STATE: {state},
						ZIPCODE: {zipcode}
				})
		`,
		Parameters: neoism.Props{"user_id": register.USER_ID,
						"first_name": register.FIRST_NAME,
						"last_name": register.LAST_NAME,
						"age": register.AGE,
						"address": register.ADDRESS,
						"township": register.TOWNSHIP,
						"state": register.STATE,
						"zipcode": register.ZIPCODE},
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return false, err
		}

		// Create Student Node
		if register.RUID != "null" { 
			fmt.Printf("Creating Student Node\n")
			cq = neoism.CypherQuery{
			Statement: `
			CREATE (n:Student {USER_ID: {user_id},
							RUID: {ruid},
							DEGREE: {degree},
							CAMPUS: {campus},
							YEAR: {year}
					})
			`,
			Parameters: neoism.Props{"user_id": register.USER_ID,
							"ruid": register.RUID,
							"degree": register.DEGREE,
							"campus": register.CAMPUS,
							"year": register.YEAR},
			}

			err = database.Cypher(&cq)
			if err != nil {
				fmt.Printf("Error Cyphering To Database\n")
				return false, err
			}
		}
		
		fmt.Printf("Successfully Registered New User\n")
		return true, nil	
	}

	return false, nil
}

func RegisterBusiness(register Util.RegisterBusiness_struct) (bool, error) {
	fmt.Printf("Accessing NEO4J Server Cyphering Business Register\n")
	Util.PrintRegisterBusiness(&register)

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err	
	}

	// Check If User Business Login Node Exists
	businessExists, err := FindBusiness(register.NAME)
	if err != nil {
		return false, err
	}

	// Create User Business Login Node
	if !businessExists {
		fmt.Printf("Creating Business User Node\n")
		cq := neoism.CypherQuery{
		Statement: `
			CREATE (n:Business {USER_ID: {user_id}, PASSWORD: {password}})
		`,
		Parameters: neoism.Props{"user_id": register.USER_ID, "password": register.PASSWORD},
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return false, err
		}

		// Create Business Node 

		fmt.Printf("Creating Business Node\n")
		cq = neoism.CypherQuery{
		Statement: `
			CREATE (n:Activity {NAME: {name},
						{ADDRESS: {address}},
						{TOWNSHIP: {township}},
						{MIN_AGE: {min_age}},
						{MAX_AGE: {max_age}},
						{MIN_PEOPLE: {min_people}},
						{MAX_PEOPLE: {max_people}},
						{MAIN_CATEGORY: {main}},
						{SUB_CATEGORY: {sub}}
				})
		`,
		Parameters: neoism.Props{"name": register.NAME,
						"address": register.ADDRESS,
						"township": register.TOWNSHIP,
						"min_age": register.MIN_AGE,
						"max_age": register.MAX_AGE,
						"min_people": register.MIN_PEOPLE,
						"max_people": register.MAX_PEOPLE,
						"main": register.MAIN_CATEGORY,
						"sub": register.SUB_CATEGORY},
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return false, err
		}
		
		fmt.Printf("Successfully Registered New Business\n")
		return true, nil
	}

	return false, nil
}

