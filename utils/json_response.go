package utils

import (
	"encoding/json"
)

// JSONResponse 回傳json data
func JSONResponse(code int, data interface{}, msg string) []byte {
	dict := make(map[string]interface{})
	dict["code"] = code
	dict["data"] = data
	dict["msg"] = msg

	res, err := json.Marshal(dict)
	if err != nil {
		return []byte(err.Error())
	}
	return res
}
