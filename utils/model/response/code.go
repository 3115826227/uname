package response

type PhoneCodeResponse struct {
	BizID     string `json:"biz_id"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}
