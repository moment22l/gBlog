package error_code

import (
	"encoding/json"
	"fmt"
	"os"
)

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 // 系统错误
)

type ErrorMap map[ErrorCode]string

func InitErrorCode() (ErrorMap, error) {
	var errorMap ErrorMap
	file, err := os.ReadFile("utils/error_code/err_code.json")
	if err != nil {
		return nil, fmt.Errorf("read err_code.json failed")
	}
	err = json.Unmarshal(file, &errorMap)
	if err != nil {
		return nil, fmt.Errorf("parse err_code.json failed")
	}
	return errorMap, nil
}
