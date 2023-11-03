package util

import (
	"encoding/json"
	"fmt"
)

func Print(data interface{}) {
	indentedJsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return
	}
	fmt.Println(string(indentedJsonData))
}
