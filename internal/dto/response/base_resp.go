package response

type Base struct {
	Data  any    `json:"data"`
	Error string `json:"error,omitempty"`
	ReqID string `json:"request_id"`
}

func NewBaseResponse(data any, err string, reqId string) Base {
	return Base{
		Data:  data,
		Error: err,
		ReqID: reqId,
	}
}
