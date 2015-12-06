package db

import
(
	"fmt"
	"github.com/jmcvetta/neoism"
	"math/rand"
	"strconv"
)

func GenerateCouponID() (string) {
	return strconv.Itoa(rand.Intn(999999999))
}

func FindCoupon(id string) (bool, error){
	database, err := neoism.Connect("http://neo4j:12345@localhost:7474/db/data")
	
	if err != nil {
		fmt.Printf("NEO4J Connection FAILED\n")
		return false, err	
	}
	
	res := []struct {
        A   string `json:"n.ID"` 
    }{}

	cq := neoism.CypherQuery{
    	Statement: `
        	MATCH (n:Coupon)
       		WHERE n.COUPON_ID = {id}
        	RETURN n.ID
    	`,
    	Parameters: neoism.Props{"id": id},
    	Result:     &res,
	}

	err = database.Cypher(&cq)
	if err != nil {
		fmt.Printf("Error Cyphering To Database\n")
		return false, err
	}

	if len(res) == 0 {
		return false, nil
	}

	return true, nil
}

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

