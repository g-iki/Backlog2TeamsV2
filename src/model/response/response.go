package response

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *Response) SetMessage(s string) {
	r.Message = s
}

func (r *Response) GetMessage() string {
	return r.Message
}

func (r *Response) SetData(d interface{}) {
	r.Data = d
}

func (r *Response) GetData() interface{} {
	return r.Data
}
