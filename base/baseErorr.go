package base

import "github.com/gin-gonic/gin"

// hanlde error
type ErrorReponse struct {
	Error string `json:"error"`
	Message string `json:"message"` 
}

type CustomeError struct {
	Message string `json:"message"`
}

var BaseError = "Product Price is requried!"


// SendError is a utility function to send a standardized error response
// gin.Context handle Http request(create update insert delete) just waiter hanlde request form customer
func SendError(handleRequest *gin.Context, statusCode int, message string, err_message error) {
    handleRequest.JSON(statusCode, ErrorReponse{
        Error:   err_message.Error(),
        Message: message,
    })
    handleRequest.Abort() // Abort use for ensuring response is sent immediately!
}

// send error 
func ErrorMessage(handleRequest *gin.Context, statusCode int ,message string){
	handleRequest.JSON(statusCode,CustomeError{
		Message: message,
	})
}

// custome success message 
func SendSuccess(handleRequest *gin.Context, statusCode int, data interface{}) {
    handleRequest.JSON(statusCode, data);
}

// custome success message 
func SendDeleteSuccess(handleRequest *gin.Context, statusCode int, message string) {
    handleRequest.JSON(statusCode,message);
}
// handle update 
// send error 
func UpdateMessage(handleRequest *gin.Context, statusCode int ,message string, data interface{}){
	handleRequest.JSON(statusCode,CustomeError{
		Message: message,
	})
}
