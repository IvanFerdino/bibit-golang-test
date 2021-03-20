package commons

import (
	"fmt"
	"net/http"
)

const INVALID_PARAMS = "Parameter tidak valid"
const SERVER_ERROR	= "Terjadi kesalahan pada server"

type CustomApiError struct {
	Result interface{} `json:"result"`
	Message  string `json:"message"`
	Code     int	`json:"code"`
}

func (e *CustomApiError) Error() string {
	return fmt.Sprintf("[%v] %v ",e.Code,e.Message)
}

func ApiSuccessResponse(result interface{}) CustomApiError {
	return CustomApiError{Result: result, Message: "Success", Code:http.StatusOK}
}

func ApiErrorResponse(message string,code int) CustomApiError {
	return CustomApiError{Result: nil,Message: message,Code:code}
}

