package buidlResp

import "encoding/json"

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func BuildResponse(code int, msg string, data interface{}) (bytes []byte, err error) {
	var (
		response Response
	)
	response.Code = code
	response.Msg = msg
	response.Data = data

	bytes, err = json.Marshal(response)
	return
}
