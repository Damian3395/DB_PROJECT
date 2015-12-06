package db

import
(
	"fmt"
	"github.com/jmcvetta/neoism"
	Util "./util"
)

func RegisterUser(register Util.RegisterUser_struct) (string, error) {
	fmt.Printf("Accessing NEO4J Server Cyphering User Register\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}

	// Check If User Login Node Exists
	userExists, err := FindUser(register.USER_ID)
	if err != nil {
		return "Username ID Not Available", err
	}

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
			return "Internal Service Error", err
		}

		// Create Consumer Node
		fmt.Printf("Creating Consumer Node\n")
		cq = neoism.CypherQuery{
		Statement: `
			CREATE (n:Consumer {USER_ID: {user_id},
						FIRST_NAME: {first_name},
						LAST_NAME: {last_name},
						AGE: {age},
						GENDER: {gender},
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
						"gender": register.GENDER,
						"address": register.ADDRESS,
						"township": register.TOWNSHIP,
						"state": register.STATE,
						"zipcode": register.ZIPCODE},
		}

		err = database.Cypher(&cq)
		if err != nil {
			fmt.Printf("Error Cyphering To Database\n")
			return "Internal Service Error", err
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
				return "Internal Service Error", err
			}
		}
		
		fmt.Printf("Successfully Registered New User\n")
		return "Success", nil	
	}

	return "Failed", nil
}

func RegisterBusiness(register Util.RegisterBusiness_struct) (string, error) {
	fmt.Printf("Accessing NEO4J Server Cyphering Business Register\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error",  err
	}

	// Check If User Business Login Node Exists
	businessExists, err := FindBusiness(register.USER_ID)
	if err != nil {
		return "Username ID Not Available", err
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
			return "Internal Service Error", err
		}

		// Create Business Node 

		fmt.Printf("Creating Business Node\n")
		cq = neoism.CypherQuery{
		Statement: `
			CREATE (n:Activity {ID: {user_id},
						NAME: {name},
						ADDRESS: {address},
						TOWNSHIP: {township},
						CAMPUS: {campus},
						MIN_AGE: {min_age},
						MAX_AGE: {max_age},
						MIN_PEOPLE: {min_people},
						MAX_PEOPLE: {max_people},
						MAIN_CATEGORY: {main},
						SUB_CATEGORY: {sub}
				})
		`,
		Parameters: neoism.Props{ "user_id": register.USER_ID,
						"name": register.NAME,
						"address": register.ADDRESS,
						"township": register.TOWNSHIP,
						"campus": register.CAMPUS,
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
			return "Internal Service Error", err
		}
		
		fmt.Printf("Successfully Registered New Business\n")
		return "Success", nil
	}

	return "Failed", nil
}

