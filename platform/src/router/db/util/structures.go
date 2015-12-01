package util

type Login_struct struct {
	ID string
	PASSWORD string
}

type RegisterUser_struct struct {
	FIRST_NAME string
	LAST_NAME string
	AGE string
	ADDRESS string
	TOWNSHIP string
	STATE string
	ZIPCODE string
	RUID string
	DEGREE string
	CAMPUS string
	YEAR string
	USER_ID string
	PASSWORD string
}

type RegisterBusiness_struct struct {
	NAME string
	ADDRESS string
	TOWNSHIP string
	MIN_AGE string
	MAX_AGE string
	MIN_PEOPLE string
	MAX_PEOPLE string
	MAIN_CATEGORY string
	SUB_CATEGORY string
	USER_ID string
	PASSWORD string
}

type Business_struct struct {
	NAME string
	ADDRESS string
	TOWNSHIP string
	MIN_AGE string
	MAX_AGE string
	MIN_PEOPLE string
	MAX_PEOPLE string
	MAIN_CATEGORY string
	SUB_CATEGORY string
}

type Consumer_struct struct{
	FIRST_NAME string
	LAST_NAME string
	AGE string
	ADDRESS string
	TOWNSHIP string
	STATE string
	ZIPCODE string
}

type Student_struct struct {
	USER_ID string
	RUID string
	DEGREE string
	CAMPUS string
	YEAR string
}
