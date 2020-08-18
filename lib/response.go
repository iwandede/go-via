package lib

type Response struct {
	ResponseID   string      `json:"response_id"`
	ResponseCode int16       `json:"code"`
	Success      bool        `json:"success"`
	Message      string      `json:"message"`
	Pagination   interface{} `json:"pagination,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        interface{} `json:"error,omitempty"`
}

func ResponseSuccess(data interface{}) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 200,
		Success:      true,
		Message:      "The request was fullfield",
		Pagination:   data,
		Data:         data,
	}

	return response
}

func ResponseNotFound(err interface{}) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 404,
		Success:      false,
		Message:      "Error data not found!",
		Error:        err,
	}

	return response
}

func ResponseUnauthorized(err interface{}) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 401,
		Success:      false,
		Message:      "Unauthorized",
		Error:        err,
	}

	return response
}

func ResponseForbidden(err interface{}) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 403,
		Success:      false,
		Message:      "Forbidden",
		Error:        err,
	}

	return response
}

func ResponseInternalError(err interface{}) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 500,
		Success:      false,
		Message:      "Internal server error",
		Error:        err,
	}

	return response
}

func ResponseConflict(err interface{}) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 409,
		Success:      false,
		Message:      "Conflict",
		Error:        err,
	}

	return response
}

func ResponseBadRequest(err interface{}) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 400,
		Success:      false,
		Message:      "Bad request!",
		Error:        err,
	}

	return response
}

func ResponseMethodNotAllowed(err interface{}) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 405,
		Success:      false,
		Message:      "Method not allowed",
		Error:        err,
	}

	return response
}

func ResponseOther(err interface{}, text string) *Response {
	response := &Response{
		ResponseID:   GenerateID(),
		ResponseCode: 500,
		Success:      false,
		Message:      text,
		Error:        err,
	}

	return response
}
