package controller

import (
	"IvanFerdino/bibit-golang-test/commons"
	"net/http"
)

func HandleRestSuccess(resp interface{}) commons.CustomApiError {
	response:=commons.ApiSuccessResponse(resp)
	return response
}

func HandleRestError(err error) (int,commons.CustomApiError) {
	var errcode int
	var errmsg string
	if _,ok:= err.(*commons.CustomApiError); ok {
		errcode=err.(*commons.CustomApiError).Code
		errmsg=err.(*commons.CustomApiError).Message
	}else{
		errcode=http.StatusInternalServerError
		errmsg=err.Error()
	}
	return errcode, commons.ApiErrorResponse(errmsg,errcode)
}
