package helper

//FormatResponse is used for static shape json return
type FormatResponse struct {
	Data interface{} `json:"data"` //dynamic
	Total string `json:"total,omitempty"`
	Response bool `json:"response"`
}

//SuccessResponse method is to inject data value to dynamic success response
func SuccessResponse(data interface{}, total string, response bool) FormatResponse {
	res := FormatResponse{
		Data:    data,
		Total:  total,
		Response: response,
	}
	return res
}
