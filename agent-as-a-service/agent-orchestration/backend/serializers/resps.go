package serializers

type Resp struct {
	Result interface{} `json:"result"`
	Data   interface{} `json:"data"`
	Error  error       `json:"error"`
	Count  *uint       `json:"count,omitempty"`
}

type APIResponse struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type RespExtra struct {
	Result      interface{} `json:"result"`
	Extra       interface{} `json:"extra"`
	Error       error       `json:"error"`
	Count       *uint       `json:"count,omitempty"`
	CountUnread *uint       `json:"count_unread,omitempty"`
}

type RespChart struct {
	Result   interface{} `json:"result"`
	Error    error       `json:"error"`
	TimeFrom uint        `json:"TimeFrom"`
	TimeTo   uint        `json:"TimeTo"`
}
