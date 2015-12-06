package db

import
(
	"fmt"
	"github.com/jmcvetta/neoism"
	util "./util"
)

func UpdateBusinessGeneral(update util.UpdateGeneralBusiness_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering UpdateBusinessGeneral\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
		A	string `json:"n.NAME"`
    }{}
	
	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Activity)
       		WHERE n.ID = {user_id}
			SET n.NAME = {name}, n.MIN_AGE = {min_age}, n.MAX_AGE = {max_age}, n.MIN_PEOPLE = {min_people},
									n.MAX_PEOPLE = {max_people}, n.MAIN_CATEGORY = {main_category}, n.SUB_CATEGORY = {sub_category}
        	RETURN n.NAME
    	`,
    	Parameters: neoism.Props{"user_id": update.ID, "name": update.NAME, "min_age": update.MIN_AGE,
									"max_age": update.MAX_AGE, "min_people": update.MIN_PEOPLE, "max_people": update.MAX_PEOPLE,
									"main_category": update.MAIN_CATEGORY, "sub_category": update.SUB_CATEGORY},
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

func UpdateBusinessAddress(update util.UpdateAddressBusiness_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering UpdateBusinessAddress\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
		A	string `json:"n.NAME"`
    }{}
	
	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Activity)
       		WHERE n.ID = {user_id}
        	SET n.ADDRESS = {address}, n.TOWNSHIP = {township}, n.CAMPUS = {campus}
			RETURN n.NAME
    	`,
    	Parameters: neoism.Props{"user_id": update.ID, "address": update.ADDRESS, "township": update.TOWNSHIP, "campus": update.CAMPUS},
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

func UpdateBusinessAll(update util.Business_struct) (string, error){
	fmt.Printf("Accessing NEO4J Server: Cyphering UpdateBusinessAll\n")

	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return "Internal Service Error", err	
	}
	
	res := []struct {
		A	string `json:"n.NAME"`
    }{}
	
	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Activity)
       		WHERE n.ID = {user_id}
        	SET n.ADDRESS = {address}, n.TOWNSHIP = {township}, n.CAMPUS = {campus}, n.MIN_AGE = {min_age},
									n.MAX_AGE = {max_age}, n.MIN_PEOPLE = {min_people},
									n.MAX_PEOPLE = {max_people}, n.MAIN_CATEGORY = {main_category}, n.SUB_CATEGORY = {sub_category}

			RETURN n.NAME
    	`,
    	Parameters: neoism.Props{"user_id": update.ID,
									"name": update.NAME, "min_age": update.MIN_AGE, "max_age": update.MAX_AGE, 
									"min_people": update.MIN_PEOPLE, "max_people": update.MAX_PEOPLE,
									"main_category": update.MAIN_CATEGORY, "sub_category": update.SUB_CATEGORY, 
									"address": update.ADDRESS, "township": update.TOWNSHIP, "campus": update.CAMPUS},
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
