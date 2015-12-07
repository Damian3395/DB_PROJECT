package util

type User_struct struct {
	ID string
}

type Login_struct struct {
	ID string
	PASSWORD string
}

type GetCoupon_struct struct{
	ID string
	COUPON_ID string
}

type Ticket_struct struct {
	ID string
	COUPON_ID string
	VALID string
}

type QuerySearch_struct struct {
	ID string
	KEY_WORDS string
	MAIN_CATEGORY string
	SUB_CATEGORY string
	CAMPUS string
	MIN_AGE string
	MAX_AGE string
	MIN_PEOPLE string
	MAX_PEOPLE string
}

type CreateCoupon_struct struct {
	ID string
	TYPE string
	DAY string
	MONTH string
	YEAR string
	VALID string
	STUDENT string
}

type RegisterUser_struct struct {
	FIRST_NAME string
	LAST_NAME string
	AGE string
	GENDER string
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

type UpdateAllUser_struct struct{
	ID string
	FIRST_NAME string
	LAST_NAME string
	AGE string
	GENDER string
	ADDRESS string
	TOWNSHIP string
	STATE string
	ZIPCODE string
	RUID string
	DEGREE string
	CAMPUS string
	YEAR string
}

type UpdateGeneralUser_struct struct{
	ID string
	FIRST_NAME string
	LAST_NAME string
	AGE string
	GENDER string
}

type UpdateAddressUser_struct struct{
	ID string
	ADDRESS string
	TOWNSHIP string
	STATE string
	ZIPCODE string
}

type UpdateGeneralBusiness_struct struct{
	ID string
	NAME string
	MIN_AGE string
	MAX_AGE string
	MIN_PEOPLE string
	MAX_PEOPLE string
	MAIN_CATEGORY string
	SUB_CATEGORY string
}

type UpdateAddressBusiness_struct struct{
	ID string
	ADDRESS string
	TOWNSHIP string
	CAMPUS string
}

type RegisterBusiness_struct struct {
	NAME string
	ADDRESS string
	TOWNSHIP string
	CAMPUS string
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
	ID string
	NAME string
	ADDRESS string
	TOWNSHIP string
	CAMPUS string
	MIN_AGE string
	MAX_AGE string
	MIN_PEOPLE string
	MAX_PEOPLE string
	MAIN_CATEGORY string
	SUB_CATEGORY string
}

type Consumer_struct struct{
	ID string
	FIRST_NAME string
	LAST_NAME string
	AGE string
	GENDER string
	ADDRESS string
	TOWNSHIP string
	STATE string
	ZIPCODE string
}

type Student_struct struct {
	ID string
	RUID string
	DEGREE string
	CAMPUS string
	YEAR string
}
