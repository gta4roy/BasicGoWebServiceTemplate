package model

import "fmt"

type ErrorType int

const (
	WRONG_INPUTS ErrorType = iota
	ERROR_IN_SAVING
	SUCCESS
)

const (
	CODE_WRONG_INPUTS    string = "10101"
	CODE_ERROR_IN_SAVING string = "10102"
	CODE_SUCCESS         string = "10103"
)

const (
	ERR_MSG_WRONG_INPUTS string = "Wrong inputs json"
	ERR_MSG_IN_SAVING    string = "Error in DB layer"
	MSG_SUCCESS_SAVE     string = "Successfull Saved the address in DB"
	MSG_UNSUCCESS_SAVE   string = "Failed in saving the address in DB"
)

type PersonModel struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (obj PersonModel) ToString() string {
	s := fmt.Sprintf(" \"name\" : %s , \"Phone\" :  %s }", obj.Name, obj.Phone)
	return s
}

type PersonModelArray struct {
	PersonRecords []PersonModel `json:"book"`
}

type ResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
