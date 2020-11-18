package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// ParserToMap 將 post request body 轉換為 map
func ParserToMap(body io.ReadCloser) map[string]interface{} {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil
	}
	dict := make(map[string]interface{})
	err = json.Unmarshal(data, &dict)
	if err != nil {
		return nil
	}
	return dict
}
