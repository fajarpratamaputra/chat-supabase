package helpers

/*
*
@desc Success response JSON

@param message string
@param data interface{}

@return interface{}
*/
func ResponseSuccess(statusCode, message string, data interface{}) interface{} {
	return map[string]interface{}{
		"data": data,
		"status": map[string]interface{}{
			"code":           statusCode,
			"message_server": message,
			"message_client": message,
		},
	}
}

/*
*
@desc Error response JSON

@param message string
@param data interface{}

@return interface{}
*/
func ResponseError(statusCode int, messageServer, messageClient string, data interface{}) interface{} {
	return map[string]interface{}{
		"data": data,
		"status": map[string]interface{}{
			"code":           statusCode,
			"message_server": messageServer,
			"message_client": messageClient,
		},
	}
}
