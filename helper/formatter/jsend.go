package formatter

type Jsend struct {
	Status  string      `json:"status" binding:"required"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildResponse(status string, msg string, data interface{}) Jsend {
	return Jsend{Status: status, Message: msg, Data: data}
}

func FailResponse(msg string) Jsend {
	return BuildResponse("failed", msg, nil)
}

func ErrorResponse(msg string) Jsend {
	return BuildResponse("error", msg, nil)
}

func SuccessResponse() Jsend {
	return BuildResponse("success", "", nil)
}

func ObjectResponse(data interface{}, envelope string) Jsend {
	if envelope != "" {
		data = map[string]interface{}{
			envelope: data,
		}
	}
	return BuildResponse("success", "", data)
}
